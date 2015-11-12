package main

import (
	"reflect"
	"./tests"
)

func main() {
	languages := []string{"Ruby", "Golang"}
	object := tests.Tests{}
	object.Languages = languages
	tests := reflect.TypeOf(object)
	for i := 0; i < tests.NumMethod(); i++ {
		method := tests.Method(i)
		if (method.Name[0:5] == "XTest") {
			object.Method = method.Name
			in := []reflect.Value{reflect.ValueOf(object)}
			method.Func.Call(in)
		}
	}
}
