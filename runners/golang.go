package runners

import (
	"../libs/sssa-golang"
	"strconv"
)

func (r *Runners) Golang_Create() {
	args, err := (r.Input).([]string)
	if !err {
		r.Output = nil
		panic("Error at Golang_Create!")
	} else {
		min, _ := strconv.Atoi(args[0])
		shares, _ := strconv.Atoi(args[1])
		secret := string(args[2])
		r.Output = sssa.Create(min, shares, secret)
	}
}

func (r *Runners) Golang_Combine() {
	args, err := (r.Input).([]string)
	if !err {
		r.Output = nil
		panic("Error at Golang_Combine!")
	} else {
		r.Output = sssa.Combine(args)
	}
}
