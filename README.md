# jsonpath
Minimalistic implementation of jsonpath

Based on http://goessner.net/articles/JsonPath/.

But a better grammar definition is here: https://info.teradata.com/htmlpubs/DB_TTU_16_00/index.html#page/Teradata_Data_Types/B035-1150-160K/pkq1472240532839.html

Install:

    go get github.com/gobs/jsonpath
    
And use it as a library:

    import "github.com/gobs/jsonpath"
    
    jsonpathexpr := "$.items[3].stuff"
    
    p := jsonpath.NewProcessor()
    err := p.Parse(jsonpathexpr)
    if err != nil {
                os.Exit(1)
    }

    // loaded from JSON
    input := map[string]interface{} {
        "name": "joe",
        "items": []interface{} {
          "zero",
          "one",
          "two",
          map[string]interface{
            "stuff": "three"
          }
     }
     
     ret := p.Process(input)
     fmt.Println(ret)

Or install the command:

    go get github.com/gobs/jsonpath/cmd/jsonpath
    
And run it:

    jsonpath "$.items[3].stuff" input.json
    
