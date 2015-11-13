package runners

import (
	"github.com/SSSaaS/go-libtest/settings"
	"io/ioutil"
	"math/rand"
	"strings"
	"encoding/json"
	"os/exec"
	"fmt"
)

type Runners struct {
	Input    interface{}
	Output   interface{}
	Settings settings.Settings
	datacode string
}

func (r *Runners) Random() {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, 15)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    r.datacode = string(b)
}

func (r *Runners) CallRunner(language settings.Langs, function string) {
	r.Random()

	for f := range r.Settings.Functions {
		if r.Settings.Functions[f].Name == function {
			for c := range r.Settings.Functions[f].Code {
				if r.Settings.Functions[f].Code[c].Type == language.Type {
					var basepath string = "./tmp/" + language.Name + "-" + language.Type + "-" + r.datacode + "-"

					var prog_data string = r.Settings.Functions[f].Code[c].Contents
					prog_data = strings.Replace(prog_data, "<INPUTFILE>", basepath + "input", -1)
					prog_data = strings.Replace(prog_data, "<OUTPUTFILE>", basepath + "output", -1)

					input, err := json.Marshal(r.Input)
					r.Check(err)

					err = ioutil.WriteFile(basepath + "prog." + language.Extension, []byte(prog_data), 0644)
					r.Check(err)

					err = ioutil.WriteFile(basepath + "input", input, 0644)
					r.Check(err)

					if language.Precommand != "" {
						args := strings.Split(language.Precommand, " ")
						comm := exec.Command(args[0], args[1:]...)
						message, err := comm.CombinedOutput()
						if (err != nil) {
							fmt.Println("Pre-Command Error:\n"+string(message)+"\n")
							fmt.Println(string(message))
							r.Check(err)
						}
					}

					if language.Command != "" {
						args := strings.Split(language.Command + " " + basepath + "prog." + language.Extension, " ")
						comm := exec.Command(args[0], args[1:]...)
						message, err := comm.CombinedOutput()
						if (err != nil) {
							fmt.Println("Command Error:\n"+string(message)+"\n")
							r.Check(err)
						}

						output_data, err := ioutil.ReadFile(basepath + "output")
						r.Check(err)

						var output interface{}
						err = json.Unmarshal(output_data, &output)
						r.Check(err)

						r.Output = output
					} else {
						fmt.Println("Empty command!?")
					}

					if language.Postcommand != "" {
						args := strings.Split(language.Postcommand, " ")
						comm := exec.Command(args[0], args[1:]...)
						message, err := comm.CombinedOutput()
						if (err != nil) {
							fmt.Println("Post-Command Error:\n"+string(message)+"\n")
							r.Check(err)
						}
					}

					return
				}
			}
		}
	}

	panic("Error! Matching language and function not found: " + function + " of " + language.Type)
}

func (r *Runners) Check(err error) {
	if err != nil {
		panic(err)
	}
}
