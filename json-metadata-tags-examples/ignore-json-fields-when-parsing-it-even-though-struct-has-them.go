// Here’s where the `json:"-"` annotation comes in handy

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy6 struct {
	Name string `json:"name"`
	Number int64 `json:"number"`
	Pointer *string `json:"-"`   // we want to ignore this during JSON parsing (unmarshalling) or writing out json (marshalling)
}

func main() {
	byt := []byte(`
		{
			"name": "oka",
			"number": 6,
			"pointer": "yes"
		}
	`)

	var temp Dummy6

	err := json.Unmarshal(byt, &temp)
	if err != nil {
		fmt.Println("An error occurred")
		os.Exit(1)
	}

	fmt.Println(temp)
	fmt.Printf("%+v\n", temp)


	// let's see the marshalling
	oka := "oka"
	tmp := Dummy6{
		Name:    "oka",
		Number:  6,
		Pointer: &oka,
	}
	byteData, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println("An error occurred2")
		os.Exit(1)
	}
	fmt.Println(string(byteData))
}