package tests

import (
	"fmt"
	"os"
	"../runners"
)

type Tests struct {
	Method string
	Languages []string
}

func (t *Tests) XChain(data interface{}, input string, output string, result interface{}) {
	r1 := runners.Runners{}
	r2 := runners.Runners{}

	for a := range t.Languages {
		for b := range t.Languages {
			r1.Input = data
			r2.Output = nil
			r1.CallRunner(t.Languages[a], "Create")

			r2.Input = r1.Output
			r2.Output = nil

			r2.CallRunner(t.Languages[b], "Combine")

			if (r2.Output != result) {
				t.XFatal(r2.Output, result, a, b)
			}
		}
	}
}

func (t Tests) Success() {
	fmt.Println("ok @ " + t.Method)
}

func (t Tests) XFatal(output interface{}, expected interface{}, a int, b int) {
	fmt.Println("fatal error @ " + t.Method + "{" + t.Languages[a] + ":" + t.Languages[b] + "}")
	fmt.Println("Excepted: ")
	fmt.Print("\t")
	fmt.Println(expected)
	fmt.Println("Received: ")
	fmt.Print("\t")
	fmt.Println(output)
	os.Exit(1)
}

func (t Tests) Fatal(output interface{}, expected interface{}) {
	fmt.Println("fatal error @ " + t.Method)
	fmt.Println("=====OUTPUT=====")
	fmt.Println(output)
	fmt.Println("=====EXPECTED=====")
	fmt.Println(expected)
	fmt.Println("=====END=====")
	os.Exit(1)
}
