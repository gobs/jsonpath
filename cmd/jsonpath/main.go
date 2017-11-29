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
	if !p.Parse(args[0]) {
		os.Exit(1)
	}

	if len(args) == 1 {
		for _, n := range p.Nodes {
			fmt.Println(n)
		}

		return
	}

	var inf *os.File

	if args[1] == "--" { // read from stdin
		inf = os.Stdin
	} else {
		f, err := os.Open(args[1])
		if err != nil {
			fatal(err)
		}

		defer f.Close()
		inf = f
	}

	j, err := simplejson.Load(inf)
	if err != nil {
		fatal(err)
	}

	ret := p.Process(j)
	fmt.Println(simplejson.MustDumpString(ret, simplejson.Indent("  ")))
}
