package main

import (
	"testing"
)

//Unit test from problem description
func TestExecute(t *testing.T) {
	input := []string{
		"glob is I",
		"prok is V",
		"pish is X",
		"tegj is L",
		"glob glob Silver is 34 Credits",
		"glob prok Gold is 57800 Credits",
		"pish pish Iron is 3910 Credits",
		"how much is pish tegj glob glob ?",
		"how many Credits is glob prok Silver ?",
		"how many Credits is glob prok Gold ?",
		"how many Credits is glob prok Iron ?",
		"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
	}

	want := []string{
		"pish tegj glob glob is 42",
		"glob prok Silver is 68 Credits",
		"glob prok Gold is 57800 Credits",
		"glob prok Iron is 782 Credits",
		"I have no idea what you are talking about",
	}

	//Init Query Struct
	var s System
	s.Init()

	//Execute Query line by line
	for _, line := range input {
		s.Translate(line)
	}

	//t.Errorf(`%v %v`, s.galaxyCredit, s.galaxyUnit)
	//Get output
	output := s.GetOutput()

	//Match the output with what we want
	for i := range want {
		if output[i] != want[i] {
			t.Errorf(`False output = "%s", want "%s"`, output[i], want[i])
		}
	}
}
