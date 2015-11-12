package runners

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"os/exec"
)

func (r *Runners) Ruby_Create() {err :=
	os.Remove("./tmp/Ruby_Create_Input.json")
	os.Remove("./tmp/Ruby_Create_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Ruby_Create_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("ruby", "./runners/ruby/sssa-create.rb")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Ruby_Create_Output.json")
	r.Check(err)

	var output []string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}

func (r *Runners) Ruby_Combine() {
	os.Remove("./tmp/Ruby_Combine_Input.json")
	os.Remove("./tmp/Ruby_Combine_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Ruby_Combine_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("ruby", "./runners/ruby/sssa-combine.rb")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Ruby_Combine_Output.json")
	r.Check(err)

	var output string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}
