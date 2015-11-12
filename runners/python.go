package runners

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"os/exec"
)

func (r *Runners) Python2_Create() {err :=
	os.Remove("./tmp/Python2_Create_Input.json")
	os.Remove("./tmp/Python2_Create_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Python2_Create_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("python2", "./runners/python2/sssa-create.py")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Python2_Create_Output.json")
	r.Check(err)

	var output []string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}

func (r *Runners) Python2_Combine() {
	os.Remove("./tmp/Python2_Combine_Input.json")
	os.Remove("./tmp/Python2_Combine_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Python2_Combine_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("python2", "./runners/python2/sssa-combine.py")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Python2_Combine_Output.json")
	r.Check(err)

	var output string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}

func (r *Runners) Python3_Create() {err :=
	os.Remove("./tmp/Python3_Create_Input.json")
	os.Remove("./tmp/Python3_Create_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Python3_Create_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("python3", "./runners/python3/sssa-create.py")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Python3_Create_Output.json")
	r.Check(err)

	var output []string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}

func (r *Runners) Python3_Combine() {
	os.Remove("./tmp/Python3_Combine_Input.json")
	os.Remove("./tmp/Python3_Combine_Output.json")

	args, err := json.Marshal(r.Input)
	r.Check(err)

	err = ioutil.WriteFile("./tmp/Python3_Combine_Input.json", args, 0644)
	r.Check(err)

	comm := exec.Command("python3", "./runners/python3/sssa-combine.py")
	err = comm.Start()
	r.Check(err)
	err = comm.Wait()
	r.Check(err)

	file, err := ioutil.ReadFile("./tmp/Python3_Combine_Output.json")
	r.Check(err)

	var output string
	err = json.Unmarshal(file, &output)
	r.Check(err)

	r.Output = output
}
