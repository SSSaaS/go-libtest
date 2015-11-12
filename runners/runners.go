package runners

import (
	"reflect"
)

type Runners struct {
	Input interface{}
	Output interface{}
}

func (r *Runners) CallRunner(language string, function string) {
	rs := reflect.TypeOf(r)
	for i := 0; i < rs.NumMethod(); i++ {
		method := rs.Method(i)
		if (method.Name == language + "_" + function) {
			in := []reflect.Value{reflect.ValueOf(r)}
			method.Func.Call(in)
		}
	}
}

func (r *Runners) Check(err error) {
	if (err != nil) {
		panic(err)
	}
}
