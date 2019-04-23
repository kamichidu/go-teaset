package teaset

import (
	_ "github.com/mjibson/esc/embed"
)

//go:generate esc -pkg assets -o internal/assets/bindata.go ./internal/template/
