// `json:"-"` once again, it works both ways

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy7 struct {
	Name string `json:"name"`
	Number int64 `json:"number"`
	Pointer *string `json:"-"`
}

func main() {
	// let's see the marshalling
	oka := "oka"
	tmp := Dummy7{
		Name:    "oka",
		Number:  7,
		Pointer: &oka,
	}
	byteData, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println("An error occurred2")
		os.Exit(1)
	}
	fmt.Println(string(byteData))
}