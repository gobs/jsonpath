package main

import "github.com/gobs/jsonpath"
import "fmt"

func main() {
	jsonpathexpr := "$.items[3].stuff"

	p := jsonpath.NewProcessor()
	err := p.Parse(jsonpathexpr)
	if err != nil {
		return
	}

	// loaded from JSON
	input := map[string]interface{}{
		"name": "joe",
		"items": []interface{}{
			"zero",
			"one",
			"two",
			map[string]interface{}{
				"stuff": "three",
			},
		},
	}

	ret := p.Process(input)
	fmt.Println(ret)
}
