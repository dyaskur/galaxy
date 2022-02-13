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

	wants := []string{
		"pish tegj glob glob is 42",
		"glob prok Silver is 68 Credits",
		"glob prok Gold is 57800 Credits",
		"glob prok Iron is 782 Credits",
		"I have no idea what you are talking about",
	}

	DoTest(input, wants, t)
}

//Unit test custom
func TestExecute2(t *testing.T) {
	input := []string{
		"prek is I",
		"vick is V",
		"pish is X",
		"tegj is L",
		"anu is C",
		"kae is D",
		"lho is M",
		"prek prek Silver is 34 Credits",
		"prek vick Gold is 57800 Credits",
		"pish pish Iron is 3910 Credits",
		"kae anu Platinum is 30000 Credits",
		"how much is pish tegj prek prek ?",
		"how much is anu kae pish vick ?",
		"how much is anu kae lho vick ?",
		"how many Credits is prek vick Silver ?",
		"how many Credits is prek vick Gold ?",
		"how many Credits is prek vick Iron ?",
		"how many Credits is prek vick Platinum ?",
		"how many Credits is prek vick vick vick vick Platinum ?",
		"how much nganu kae?",
	}

	wants := []string{
		"pish tegj prek prek is 42",
		"anu kae pish vick is 415",
		"anu kae lho vick is 1405",
		"prek vick Silver is 68 Credits",
		"prek vick Gold is 57800 Credits",
		"prek vick Iron is 782 Credits",
		"prek vick Platinum is 200 Credits",
		"Your units are not valid",
		"I have no idea what you are talking about",
	}

	DoTest(input, wants, t)
}

//Unit test custom
func TestExecute3(t *testing.T) {
	input := []string{
		"poli is I",
		"tali is V",
		"deli is X",
		"tegj is L",
		"seli is C",
		"kuli is D",
		"peti is M",
		"poli poli Uranium is 68 Credits",
		"poli tali Vibranium is 128000 Credits",
		"deli deli Plutonium is 87000 Credits",
		"kuli seli Thorium is 15000 Credits",
		"how much is deli tegj poli poli ?",
		"how much is seli kuli deli tali ?",
		"how much is seli kuli peti tali ?",
		"how many Credits is poli tali Uranium ?",
		"how many Credits is poli tali Vibranium ?",
		"how many Credits is poli tali Plutonium ?",
		"how many Credits is poli tali Thorium ?",
		"how many Credits is poli poli poli poli tali Thorium ?",
		"how much ngseli kuli?",
	}

	wants := []string{
		"deli tegj poli poli is 42",
		"seli kuli deli tali is 415",
		"seli kuli peti tali is 1405",
		"poli tali Uranium is 136 Credits",
		"poli tali Vibranium is 128000 Credits",
		"poli tali Plutonium is 17400 Credits",
		"poli tali Thorium is 100 Credits",
		"Your units are not valid",
		"I have no idea what you are talking about",
	}
	DoTest(input, wants, t)
}

func DoTest(input []string, wants []string, t *testing.T) {
	//Init Query Struct
	var s System
	s.Init()

	//Translate input to action command
	for _, line := range input {
		s.Translate(line)
	}
	//Get output
	output := s.output

	//Match the output with what we want
	for i := range wants {
		if output[i] != wants[i] {
			t.Errorf(`False output = "%s", want "%s"`, output[i], wants[i])
		}
	}
}

//
