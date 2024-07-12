package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/berquerant/jsonast"
)

const usage = `jsonast - build json AST

Usage:

  jsonast [flags] [FILE]

This command converts json to AST but does not convert the value type of the node.
Node values ​​are always strings.

Flags:`

func Usage() {
	fmt.Fprintln(os.Stderr, usage)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	file := flag.Arg(0)

	if err := run(file); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run(file string) error {
	var r io.Reader = os.Stdin
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		r = f
		defer f.Close()
	}

	root, err := jsonast.Parse(r)
	if err != nil {
		return err
	}
	b, err := json.Marshal(root)
	if err != nil {
		return err
	}
	fmt.Printf("%s", b)
	return nil
}
