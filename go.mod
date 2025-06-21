module github.com/berquerant/jsonast

go 1.24.4

require (
	github.com/alecthomas/assert/v2 v2.11.0
	github.com/alecthomas/participle/v2 v2.1.4
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/alecthomas/repr v0.4.0 // indirect
	github.com/berquerant/marker v0.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hexops/gotextdiff v1.0.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/telemetry v0.0.0-20240522233618-39ace7a40ae7 // indirect
	golang.org/x/tools v0.29.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

tool (
	github.com/berquerant/marker
	golang.org/x/vuln/cmd/govulncheck
)
