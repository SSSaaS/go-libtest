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
	settings_file, err := ioutil.ReadFile("libtest.json")
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
	fmt.Println("total:", fmt.Sprintf("%.3f", duration.Seconds()), "s")
	os.RemoveAll("./tmp")
}
