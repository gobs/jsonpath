package main

import "github.com/gobs/jsonpath"
import "fmt"

func main() {
	jsonpathexpr := "$.items[3].stuff"

	p := jsonpath.NewProcessor()
	if !p.Parse(jsonpathexpr) {
		return
	}

	// loaded from JSON
	input := map[string]any{
		"name": "joe",
		"items": []any{
			"zero",
			"one",
			"two",
			map[string]any{
				"stuff": "three",
			},
		},
	}

	ret := p.Process(input, jsonpath.Default)
	fmt.Println(ret)
}
