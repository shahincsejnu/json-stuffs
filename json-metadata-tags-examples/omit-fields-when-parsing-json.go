// if we do not want to all of the json fields during converting it to (unmarshalling) struct
// then we can just convert that json into a struct which has our necessary fields only
// trick : we can parse the same JSON using multiple structs, to adjust different contents
// during unmarshalling from json to struct only existing struct's fields get the value fron json, rest get omitted

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy5 struct {
	Name string `json:"name"`
	// we do not need any other fields from the json during parsing, that is why we only used one field in this struct
}

func main() {
	byt := []byte(`
		{
			"name": "oka",
			"number": 5,
			"pointer": "yes"
		}
	`)

	var temp Dummy5

	err := json.Unmarshal(byt, &temp)
	if err != nil {
		fmt.Println("An error occurred")
		os.Exit(1)
	}

	fmt.Println(temp)
	fmt.Printf("%+v\n", temp)

	// we can see the output : just one field, nothing else got used
}
