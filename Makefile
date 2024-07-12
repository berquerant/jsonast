GOMOD = go mod
GOBUILD = go build -trimpath -v
GOTEST = go test -v -cover -race

ROOT = $(shell git rev-parse --show-toplevel)
BIN = dist/jsonast
CMD = "./cmd/jsonast"

.PHONY: $(BIN)
$(BIN):
	$(GOBUILD) -o $@ $(CMD)

.PHONY: test
test:
	$(GOTEST) ./...

testdata/output.json: testdata/input.json
	go run $(CMD) $< > $@

.PHONY: init
init:
	$(GOMOD) tidy

.PHONY: vuln
vuln:
	go run golang.org/x/vuln/cmd/govulncheck ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: regenarate
regenarate: clean generate

.PHONY: generate
generate: go-generate

.PHONY: clean
clean: clean-go-generate

.PHONY: go-regenerate
go-regenerate: clean-go-generate go-generate

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: clean-go-generate
clean-go-generate:
	find $(ROOT) -name "*_generated.go" -type f -delete
