package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonString :=
		`[
			  {
				"Name" : "Shahin",
				"Role" : "Software Engineer",
				"ID" : 7,
				"Male" : true,
				"Skills" : ["Docker", "K8s", "Terraform"],
				"Profiles" : {
				  "LinkedIn" : "okaoka"
				}
			  },
			
			  {
				"Name" : "Oka",
				"Skills" : ["Docker", true, "Terraform"],
				"ID" : null
			  }
	]`

	fmt.Println(jsonString)

	jsn, err := json.Marshal(jsonString)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(jsn)
}