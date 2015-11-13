package main

import (
	"encoding/json"
	"github.com/SSSaaS/go-libtest/settings"
	"github.com/SSSaaS/go-libtest/tests"
	"io/ioutil"
	"time"
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) == 1) {
		fmt.Println("Error! Must specify at least one spec file:")
		fmt.Println("")
		fmt.Println("Usage: go-libtest spec1.json [spec2.json ... specn.json]")
	}

	for i := range os.Args[1:] {
		if _, err := os.Stat(os.Args[i+1]); os.IsNotExist(err) {
			fmt.Printf("no such file or directory: %s", os.Args[i+1])
			return
		}

		fmt.Println("loading specs: " + os.Args[i+1])

		settings_file, err := ioutil.ReadFile(os.Args[i+1])
		if err != nil {
			panic(err)
		}

		os.RemoveAll("./tmp")
		os.Mkdir("./tmp", 0777)

		var settings_struct settings.Settings
		err = json.Unmarshal(settings_file, &settings_struct)
		if err != nil {
			panic(err)
		}

		t1 := time.Now()
		for t := range settings_struct.Tests {
			test := settings_struct.Tests[t]
			tester := tests.Tests{test.Method, settings_struct}
			if test.Type == "XChain" {
				tester.XChain(test.Data, test.Input, test.Output, test.Result)
			}
		}

		duration := time.Since(t1)
		fmt.Println("------")
		fmt.Println("total:", fmt.Sprintf("%.3f", duration.Seconds()), "s\n")
		os.RemoveAll("./tmp")
	}
}
