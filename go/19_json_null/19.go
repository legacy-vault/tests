package main // 19.go.

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {

	var ba []byte
	var err error
	var obj interface{}
	var str_1 string
	//var str_2 string
	var objType reflect.Type
	var objValue reflect.Value

	// Prepare some Data.
	str_1 = `
{
	"Name":null
}`
	//str_2 = `{null}` =>
	// Error:
	// invalid character 'n' looking for beginning of object key string.
	ba = []byte(str_1)
	//ba = []byte(str_2)

	// Decode to an empty Interface.
	err = json.Unmarshal(ba, &obj)
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Printf("%+v.\r\n", obj)
	objType = reflect.TypeOf(obj)
	objValue = reflect.ValueOf(obj)
	fmt.Println(objType)
	fmt.Println(objValue)

	var m map[string]interface{}
	var name interface{}
	var ok bool
	m, ok = obj.(map[string]interface{})
	if !ok {
		log.Fatal(err)
	}
	name = m["Name"]
	fmt.Println(reflect.TypeOf(name))        //
	fmt.Println(reflect.TypeOf(name).Kind()) //
}
