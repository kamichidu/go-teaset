package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kamichidu/go-teaset/internal/generator"
	"golang.org/x/tools/go/ast/astutil"
)

var (
	possibleBaseImpls = []string{
		"HashSet",
		"TreeSet",
	}
)

func usageFunc(flgs *flag.FlagSet) func() {
	s := `Usage:
  teaset-gen [options]

Mandatory options:
  -o filename    Output file.
  -base          The base implementation. (Choices: ` + strings.Join(possibleBaseImpls, ", ") + `)

Available options:
  -pkg           Destination package name.
  -set-type      The generated type name.
  -element-type  The element type of generated set.

Example:
  teaset-gen -o uuidset.go -pkg testing -base HashSet -set-type UUIDSet -element-type "github.com/gofrs/uuid".UUID
  teaset-gen -o stringset.go -pkg testing -base HashSet -set-type StringSet -element-type string
`
	return func() {
		fmt.Fprint(flgs.Output(), s)
	}
}

func run(stderr io.Writer, args []string) int {
	log.SetOutput(stderr)

	flgs := flag.NewFlagSet("teaset-gen", flag.ContinueOnError)
	flgs.SetOutput(ioutil.Discard)
	flgs.Usage = usageFunc(flgs)
	var (
		debug    = flgs.Bool("debug", false, "")
		out      = flgs.String("o", "", "")
		pkg      = flgs.String("pkg", "", "")
		eleTyp   = flgs.String("element-type", "", "")
		setTyp   = flgs.String("set-type", "", "")
		baseImpl = flgs.String("base", "", "")
	)
	if err := flgs.Parse(args[1:]); err == flag.ErrHelp {
		flgs.SetOutput(stderr)
		flgs.Usage()
		return 0
	} else if err != nil {
		fmt.Fprintln(stderr, err)
		return 128
	}
	if *out == "" || *baseImpl == "" {
		flgs.SetOutput(stderr)
		flgs.Usage()
		return 128
	}

	if *pkg == "" {
		abs, err := filepath.Abs(*out)
		if err != nil {
			panic(err)
		}
		*pkg = filepath.Base(filepath.Dir(abs))
	}
	log.Printf("Output package %q ...", *pkg)
	if *eleTyp == "" {
		*eleTyp = "interface{}"
	}
	log.Printf("Output element type %q ...", *eleTyp)

	if *debug {
		log.Print("Debug mode enabled ...")
		generator.UseLocal = true
		generator.Debug = true
	}

	var w io.Writer
	if *out == "-" {
		w = os.Stdout
	} else {
		dir := filepath.Dir(*out)
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("failed to create dirs %q: %v", dir, err)
			return 1
		}
		f, err := os.Create(*out)
		if err != nil {
			log.Printf("failed to create file %q: %v", *out, err)
			return 1
		}
		defer f.Close()
		w = f
	}

	log.Printf("Parsing %q sources ...", *baseImpl)
	fset, aFile, err := generator.ParseFile(*baseImpl)
	if err != nil {
		switch err := err.(type) {
		case scanner.ErrorList:
			for _, e := range err {
				log.Print(e)
			}
		default:
			log.Print(err)
		}
		return 1
	}
	eleTypPkgPath, eleTypPkgName, eleTypName := generator.ParseElementType(*eleTyp)
	var imports []generator.Import
	if eleTypPkgPath != "" {
		imports = append(imports, generator.Import{
			PkgName: eleTypPkgName,
			Path:    eleTypPkgPath,
		})
	}
	var applyer []astutil.ApplyFunc
	if *debug {
		applyer = append(applyer, generator.PrintCursorInfo)
	}
	// replace package name
	applyer = append(applyer, generator.ReplaceExactIdentName("template", *pkg))
	// replace Element
	if eleTypPkgName != "" {
		applyer = append(applyer, generator.ReplaceExactIdentName("Element", eleTypPkgName+"."+eleTypName))
	} else {
		applyer = append(applyer, generator.ReplaceExactIdentName("Element", eleTypName))
	}
	// replace type name
	if *setTyp != "" {
		applyer = append(applyer, generator.ReplaceIdentName(*baseImpl, *setTyp))
	}
	log.Print("Generating ...")
	fmt.Fprintf(w, "// Code generated by %s %s; DO NOT EDIT THIS FILE\n\n", flgs.Name(), strings.Join(args[1:], " "))
	err = generator.Traverse(w, fset, aFile, imports, generator.ApplyFunc(applyer...))
	if err != nil {
		log.Print(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Stderr, os.Args))
}
