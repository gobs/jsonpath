//go:generate antlr4 -Dlanguage=Go -package parser -o parser Jsonpath.g4
package main

import (
	"fmt"
	"os"

	"github.com/gobs/jsonpath"
	"github.com/gobs/simplejson"
)

func fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fatal("usage:", os.Args[0], "$.json.path [filename]")
	}

	p := jsonpath.NewProcessor()
	err := p.Parse(args[0])

	if err != nil {
		os.Exit(1)
	}

	if len(args) == 1 {
		for _, n := range p.Nodes {
			fmt.Println(n)
		}

		return
	}

	f, err := os.Open(args[1])
	if err != nil {
		fatal(err)
	}

	defer f.Close()

	j, err := simplejson.Load(f)
	if err != nil {
		fatal(err)
	}

	ret := p.Process(j)
	fmt.Println(simplejson.MustDumpString(ret, simplejson.Indent("  ")))
}
