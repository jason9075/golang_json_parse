package main

import (
	"encoding/json"
	"fmt"
)

type DataStruct struct {
	Field1 string
	Field2 int
	Field3 string
}

func (data *DataStruct) UnmarshalJSON(b []byte) error {

	var i interface{}
	json.Unmarshal(b, &i)

	top := i.(map[string]interface{})

	insideMap := top["inside"]
	inside := insideMap.(map[string]interface{})

	data.Field1 = inside["field1"].(string)
	data.Field2 = int(inside["field2"].(float64))
	data.Field3 = top["field3"].(string)

	return nil
}

func main() {

	json_message := `
    { 
        "inside" : { 
            "field1" : "111", 
            "field2" : 2 , 
            "field_xxx" : "xxx" 
        }, 
        "field_yyy" : "yyy",
        "field3" : "field3"
    }`
	data := &DataStruct{}
	json.Unmarshal([]byte(json_message), &data)

	fmt.Printf("%+v\n", data)
}
