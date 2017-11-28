# jsonpath
Minimalistic implementation of jsonpath, based on http://goessner.net/articles/JsonPath/.

But a better grammar definition is [here](https://info.teradata.com/htmlpubs/DB_TTU_16_00/index.html#page/Teradata_Data_Types/B035-1150-160K/pkq1472240532839.html).

## Documentation

https://godoc.org/github.com/gobs/jsonpath

## Usage

Install:

    go get github.com/gobs/jsonpath
    
And use it as a library:

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

Or install the command:

    go get github.com/gobs/jsonpath/cmd/jsonpath
    
And run it:

    jsonpath "$.items[3].stuff" input.json
    
