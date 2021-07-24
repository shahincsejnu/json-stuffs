# json-stuffs

This repository contains the stuffs related to JSON (or JSON in Golang to be specific).

## JSON

- JavaScript Object Notation
- Data representation format
- Commonly used for APIs and Configs
- Lightweight and easy to read/write
- It's a superset of JavaScript
- Integrates easily with most languages
- It's extension is `.json`

## JSON (Data) Types

- It supports `BASONN` types of data :
    - `B` : Booleans, Ex: true, false etc.
    - `A` : Arrays, Ex: [1, 2, 3], ["Hello", "World"] etc.
    - `S` : Strings, Ex: "Hello World", "Shahin" etc.
    - `O` : Objects, Ex: {"key" : <value>}, {"age": 25} etc.
    - `N` : Numbers, Ex: 1, 2, 1.5, -30, -3.5, 1.2e10 etc any time of numbers
    - `N` : Null
    
## Syntax & Rules

- Need to provide `,` after the `"key": <value>,` pair (in json object) except the last one
  - Ex:
    ```json
      {
        "Name" : "Shahin",
        "Role" : "Software Engineer",
        "ID" : 7,
        "Male" : true
      }
    ```
    
- JSON Objects can be nested, like : 
  - Ex:
    ```json
      {
        "Name" : "Shahin",
        "Role" : "Software Engineer",
        "ID" : 7,
        "Male" : true,
        "Profiles" : {
          "LinkedIn" : "okaoka"
         }
    }
    ```

- JSON Objects can be represented in arrays, like :
  - Ex:
    ```json
      [
        {
          "Name" : "Shahin",
          "Role" : "Software Engineer",
          "ID" : 7,
          "Male" : true,
          "Profiles" : {
            "LinkedIn" : "okaoka"
          }
        },
        
        {
          "Name" : "Oka"
        }
      ]
    ```
  - In this case also need to put `,` after every (parallel level) objects except the last one like the example above
  
- In a `"Key": <value>` pair of json object, the value can be any valid json data types type value, like:
  - Ex:
    ```json
      {
        "Name" : "Shahin",
        "Role" : "Software Engineer",
        "ID" : 7,
        "Male" : true,
        "Skills" : ["Docker", "K8s", "Terraform"],
        "Profiles" : {
          "LinkedIn" : "okaoka"
        }
      }
    ```
    
- In an array we can put any types of values together, all values can be one type or multiple values can be put together, arrays represent by `[]` :
  - Ex:
    ```json
    [
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
        "Skills" : ["Docker", true, "Terraform"]
      }
    ]
    ```
- If in a `"key": <value>` pair of json object any field doesn't have any value then we can use `null` as the value, like:
  - Ex:
    ```json
      {
        "Name" : "Oka",
        "Skills" : ["Docker", true, "Terraform"],
        "ID" : null
      }
    ```

- If we want to use the json object as a string (json string), we need to surround the json object with back tick(``), like:
  - Ex:
    ```
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
    ```
  - See `main.go` to know how to use this json string, how to parse (Marshal and Unmarshal in Go) them to a json (from json string to json) etc.
  

## JSON in Go




# Resources 

- [Learn JSON in 10 Minutes](https://www.youtube.com/watch?v=iiADhChRriM)
- [Working with JSON in Go - Tutorial](https://www.youtube.com/watch?v=Osm5SCw6gPU) 

