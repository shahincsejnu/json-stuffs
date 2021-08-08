// `json:"omitempty"` to the rescue
// we only want to ignore the field of the struct if it's empty
// Zero-values and nil pointers are considered empty, and this can feel counterintuitive sometimes. If you want to output a 0, you better work with something like *int64, if you really want to use omitempty.
// see the Dummy9, if we use *string or *int etc then if we provider "" as string and 0 as int then they will be considered as a string and 0
// but if we use string or int then the default value of that type (0 for int, "" for string et) will considered as nil value, so at that time distinguishing between them is difficult
// because if we use pointer (*string) then we need to explicitly given `nil` as the nil value

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy8 struct {
	Name    string  `json:"name,omitempty"`
	Number  int     `json:"number,omitempty"`
	Pointer *string `json:"pointer,omitempty"`
}

type Dummy9 struct {
	Name    *string `json:"name,omitempty"`
	Number  *int    `json:"number,omitempty"`
	Pointer *string `json:"pointer,omitempty"`
}

func main() {
	tmp0 := Dummy8{
		Name:    "",
		Number:  0,
		Pointer: nil,
	}
	byt0, err := json.Marshal(tmp0)
	if err != nil {
		fmt.Print("An error occurred0")
		os.Exit(1)
	}
	fmt.Println(string(byt0))

	tmp1 := Dummy8{
		Name:    "",
		Number:  5,
		Pointer: nil,
	}
	byt1, err := json.Marshal(tmp1)
	if err != nil {
		fmt.Println("An error occurred1")
		os.Exit(1)
	}
	fmt.Println(string(byt1))

	tmp2 := Dummy9{
		Name:    nil,
		Number:  nil,
		Pointer: nil,
	}
	byt2, err := json.Marshal(tmp2)
	if err != nil {
		fmt.Println("An error occurred2")
		os.Exit(1)
	}
	fmt.Println(string(byt2))

	val := 8
	tmp3 := Dummy9{
		Name:    nil,
		Number:  &val,
		Pointer: nil,
	}
	byt3, err := json.Marshal(tmp3)
	if err != nil {
		fmt.Println("An error occurred3")
		os.Exit(1)
	}
	fmt.Println(string(byt3))

	oka := "okaoka"
	tmp4 := Dummy8{
		Name:    "oka",
		Number:  val,
		Pointer: &oka,
	}
	byt4, err := json.Marshal(tmp4)
	if err != nil {
		fmt.Println("An error occurred4")
		os.Exit(1)
	}
	fmt.Println(string(byt4))

	byt5 := []byte(`
		{
			"name": "",
			"number": 0,
			"pointer": null
		}
	`)
	// we can give `null` if a field doesn't have any value in the json
	var tmp5 Dummy8
	var tmp6 Dummy9

	err = json.Unmarshal(byt5, &tmp5)
	if err != nil {
		fmt.Println("An error occurred5")
		os.Exit(1)
	}
	fmt.Printf("%+v\n", tmp5)

	err = json.Unmarshal(byt5, &tmp6)
	if err != nil {
		fmt.Println("An error occurred6")
		os.Exit(1)
	}
	fmt.Printf("%+v\n", tmp6)
}
