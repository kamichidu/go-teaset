.PHONY: help
help:  ## show this help
	@grep -E '^[a-zA-Z_\/-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: generate  ## install teaset-gen
	go install ./cmd/teaset-gen/

.PHONY: generate
generate:  ## go generate
	go generate ./...

.PHONY: defaults
defaults:  ## generate for root package
	teaset-gen -pkg teaset -o hashset.go -base HashSet
	teaset-gen -pkg teaset -o treeset.go -base TreeSet
