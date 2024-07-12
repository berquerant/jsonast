//go:build tools
// +build tools

package main

import (
	_ "github.com/berquerant/marker"
	_ "golang.org/x/vuln/cmd/govulncheck"
)
