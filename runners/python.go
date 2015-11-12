package runners

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"os/exec"
)

func (r *Runners) Python_Create() {err :=
	os.Remove("./tmp/Python_Create_Input.json")
	os.Remove("./tmp/Python_Create_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Python_Create_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("python", "./runners/python/sssa-create.py")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Python_Create_Output.json")
	r.Check(err)

	var output []string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}

func (r *Runners) Python_Combine() {
	os.Remove("./tmp/Python_Combine_Input.json")
	os.Remove("./tmp/Python_Combine_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Python_Combine_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("python", "./runners/python/sssa-combine.py")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Python_Combine_Output.json")
	r.Check(err)

	var output string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}
