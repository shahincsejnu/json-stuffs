// let's do json unmarshalling : convert json (byte data) to built-in or custom data (here in struct)

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy3 struct {
	Name string
	Number int64
	Pointer *string
}

type Dummy4 struct {
	Name string `json:"name"`
	Number int64 `json:"number"`
	Pointer *string `json:"pointer"`
}

func main() {
	//dum3 := Dummy3{}
	//dum4 := Dummy4{}
	var dum3 Dummy3
	var dum4 Dummy4

	byt := []byte(`
		{
			"name": "oka",
			"number": 3,
			"pointer": "nothing"
		}
	`)

	err := json.Unmarshal(byt, &dum3)
	if err != nil {
		fmt.Println("An error occurred dum3")
		os.Exit(1)
	}
	fmt.Println(dum3)
	fmt.Printf("%+v\n", dum3)

	err = json.Unmarshal(byt, &dum4)
	if err != nil {
		fmt.Println("An error occurred dum4")
		os.Exit(1)
	}
	fmt.Println(dum4)
	fmt.Printf("%+v\n", dum4)

	// so we can see the output : in both dum3 and dum4 the field names is as like the struct field name (capital letter), so json tags custom field names used when converting to json (marshal)
	// also, all fields got used in this unmarshalling
}