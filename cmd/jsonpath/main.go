//go:generate antlr4 -Dlanguage=Go -package parser -o parser Jsonpath.g4
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gobs/jsonpath"
	"github.com/gobs/simplejson"
)

func fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	var options jsonpath.ProcessOptions

	enhanced := flag.Bool("e", false, "enhanced: returns key/value pairs for objects")
	collapse := flag.Bool("c", false, "collapse: if result is an array of 1 element, return the element")
	flat := flag.Bool("f", false, "flat: flatten content of results from nested dictionaries")
	flag.Parse()

	if *enhanced {
		options |= jsonpath.Enhanced
	}
	if *collapse {
		options |= jsonpath.Collapse
	}
	if *flat {
		options |= jsonpath.Flat
	}

	if flag.NArg() == 0 || flag.NArg() > 2 {
		fatal("usage:", os.Args[0], "$.json.path [filename]")
	}

	query := flag.Arg(0)
	if !strings.Contains(query, `'`) {
		query = strings.ReplaceAll(query, `"`, `'`)
	}

	p := jsonpath.NewProcessor()
	if !p.Parse(query) {
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

	ret := p.Process(j, options)
	fmt.Println(simplejson.MustDumpString(ret, simplejson.Indent("  ")))
}
