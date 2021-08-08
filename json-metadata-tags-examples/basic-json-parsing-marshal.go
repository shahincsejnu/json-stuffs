// we can do json marshal from built-in or custom data types
// here we will see two types : one with providing json tags (custom json fields names)
// another without providing custom json fields names
// json marshalling and unmarshalling only happens with exported fields
// so the struct's fields name should be in exported format (start with capital letter)

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy1 struct {
	Name    string
	Number  int64
	Pointer *string
}

type Dummy2 struct {
	Name    string  `json:"name"`
	Number  int64   `json:"number"`
	Pointer *string `json:"pointer"`
}

func main() {
	oka := "okaoka"

	dum1 := &Dummy1{
		Name:    "Oka",
		Number:  0,
		Pointer: &oka,
	}

	byt, err := json.Marshal(dum1)
	if err != nil {
		fmt.Println("An error occured")
		os.Exit(1)
	}

	fmt.Println(string(byt)) // we can see the output : json fields name is as like the struct fields names

	// we can give custom json fields/keys names
	// for that we need to use json tags in the struct, like Dummy2
	dum2 := &Dummy2{
		Name:    "Oka",
		Number:  2,
		Pointer: &oka,
	}

	byt2, err := json.Marshal(dum2)
	if err != nil {
		fmt.Println("An error occured2")
		os.Exit(1)
	}

	fmt.Println(string(byt2)) // we can see the output : our given custom field name has been used as json fields
}
