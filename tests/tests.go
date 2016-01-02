package tests

import (
	"fmt"
	"github.com/SSSaaS/go-libtest/runners"
	"github.com/SSSaaS/go-libtest/settings"
	"os"
	"time"
)

type Tests struct {
	Method   string
	Settings settings.Settings
}

func (t *Tests) XChain(data interface{}, input string, output string, result interface{}) {
	t1 := time.Now()
	r1 := runners.Runners{}
	r1.Settings = t.Settings
	r2 := runners.Runners{}
	r2.Settings = t.Settings

	for a := range t.Settings.Languages {
		for b := range t.Settings.Languages {
			r1.Input = data
			r2.Output = nil
			r1.CallRunner(t.Settings.Languages[a], input)

			r2.Input = r1.Output
			r2.Output = nil

			r2.CallRunner(t.Settings.Languages[b], output)

			if r2.Output != result {
				t.XFatal(r2.Output, result, a, b)
			}
		}
	}

	duration := time.Since(t1)
	t.Success(duration)
}

func (t Tests) Success(duration time.Duration) {
	fmt.Println("ok @ "+t.Method+" in", fmt.Sprintf("%.3f", duration.Seconds()), "s")
}

func (t Tests) XFatal(output interface{}, expected interface{}, a int, b int) {
	fmt.Println("fatal error @ " + t.Method + "{" + t.Settings.Languages[a].Name + ":" + t.Settings.Languages[b].Name + "}")
	fmt.Println("Excepted: ")
	fmt.Println(expected)
	fmt.Println("\nReceived: ")
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
