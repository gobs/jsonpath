//go:generate antlr4 -Dlanguage=Go -package parser -o parser Jsonpath.g4
package main

import (
	"flag"
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
	var enhanced bool

	flag.BoolVar(&enhanced, "e", false, "enhanced: returns key/value pairs for objects")
	flag.Parse()

	if flag.NArg() == 0 || flag.NArg() > 2 {
		fatal("usage:", os.Args[0], "$.json.path [filename]")
	}

	p := jsonpath.NewProcessor()
	if !p.Parse(flag.Arg(0)) {
		os.Exit(1)
	}

	if flag.NArg() == 1 {
		for _, n := range p.Nodes {
			fmt.Println(n)
		}

		return
	}

	var inf *os.File

	if flag.Arg(1) == "--" { // read from stdin
		inf = os.Stdin
	} else {
		f, err := os.Open(flag.Arg(1))
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

	ret := p.Process(j, enhanced)
	fmt.Println(simplejson.MustDumpString(ret, simplejson.Indent("  ")))
}
