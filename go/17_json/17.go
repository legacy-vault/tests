package main // 17.go.

import (
	"encoding/json"
	"fmt"
	"reflect"
	"log"
)

// Reference Structure.
type Data_1 struct {
	Person struct {
		Name string
		Age uint8
	}
	Hardware struct {
		Vendor string
		Year uint16
	}
}

// Structure with additional Field.
type Data_2 struct {
	Person struct {
		Name string
		Nickname string // <--
		Age uint8
	}
	Hardware struct {
		Vendor string
		Year uint16
	}
}

// Structure with missing Field.
type Data_3 struct {
	Person struct {
		Name string
		//Age uint8 // <--
	}
	Hardware struct {
		Vendor string
		Year uint16
	}
}

func main() {

	var ba []byte
	var err error
	var obj interface{}
	var mapKey reflect.Value
	var mapKeys []reflect.Value
	var str string
	var objType reflect.Type
	var objValue reflect.Value
	var refMap map[string]interface{}
	var outObj_1 Data_1
	var outObj_2 Data_2
	var outObj_3 Data_3

	// Prepare some Data.
	str = `
{
	"Person":
		{"Name":"Василий","Age":25},
	"Hardware":
		{"Vendor":"AMD","Year":2012}
}`
	ba = []byte(str)

	// 1.1. Decode to a specialized Structure 'Data_1'.
	err = json.Unmarshal(ba, &outObj_1)
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Printf("%+v.\r\n", outObj_1)
	objType = reflect.TypeOf(outObj_1)
	objValue = reflect.ValueOf(outObj_1)
	fmt.Println(objType)
	fmt.Println(objValue)
	//
	if objType != reflect.TypeOf(outObj_1){
		log.Fatal("Type is wrong!")
	} else {
		fmt.Println("Type is good.")
	}

	// 1.2. Decode to a specialized Structure 'Data_2'.
	err = json.Unmarshal(ba, &outObj_2)
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Printf("%+v.\r\n", outObj_2)
	objType = reflect.TypeOf(outObj_2)
	objValue = reflect.ValueOf(outObj_2)
	fmt.Println(objType)
	fmt.Println(objValue)
	//
	if objType != reflect.TypeOf(outObj_2){
		log.Fatal("Type is wrong!")
	} else {
		fmt.Println("Type is good.")
	}

	// 1.3. Decode to a specialized Structure 'Data_3'.
	err = json.Unmarshal(ba, &outObj_3)
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Printf("%+v.\r\n", outObj_3)
	objType = reflect.TypeOf(outObj_3)
	objValue = reflect.ValueOf(outObj_3)
	fmt.Println(objType)
	fmt.Println(objValue)
	//
	if objType != reflect.TypeOf(outObj_3){
		log.Fatal("Type is wrong!")
	} else {
		fmt.Println("Type is good.")
	}

	// 2. Decode to an empty Interface.
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
	//
	if objType != reflect.TypeOf(refMap){
		log.Fatal("Type is wrong!")
	} else {
		fmt.Println("Type is good.")
	}

	mapKeys = objValue.MapKeys()
	for _,mapKey = range mapKeys {
		fmt.Println(mapKey)
	}

	/*
	if obj.Name != nil {
		fmt.Println("Field 'Name' exists.")
	}
	*/
}
