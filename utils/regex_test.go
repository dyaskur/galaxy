package utils

import (
	"testing"
)

//test case to define variables from set unit command
func TestMatchInput(t *testing.T) {

	input := "abc is I"
	regex := `(?P<galaxyUnit>[a-z]+) is (?P<romanNum>[IVXLCDM]+)`
	wants := make(map[string]string)
	wants["galaxyUnit"] = "abc"
	wants["romanNum"] = "I"

	DoTest(input, regex, wants, t)

}

//test case to define variables from set commodity command
func TestMatchInputSetCommodityCredit(t *testing.T) {
	input := "glob glob Silver is 34 credit"
	regex := `(?P<galaxyUnits>[a-z ]+) (?P<commodity>[A-Z][a-z]*) is (?P<credit>[0-9]+) credit`
	wants := make(map[string]string)
	wants["galaxyUnits"] = "glob glob"
	wants["commodity"] = "Silver"
	wants["credit"] = "34"

	DoTest(input, regex, wants, t)

}

//test case to Define variables from how much question
func TestMatchInputGetUnits(t *testing.T) {
	input := "how much is pish tegj glob glob ?"
	regex := `how much is (?P<galaxyUnits>[a-z ]+) \?`
	wants := make(map[string]string)
	wants["galaxyUnits"] = "pish tegj glob glob"

	DoTest(input, regex, wants, t)
}

//test case to Define variables from how many question
func TestMatchInputGetCommodity(t *testing.T) {
	input := "how many credit is glob prok Silver ?"
	regex := `how many credit is (?P<galaxyUnits>[a-z ]+) (?P<commodity>[A-Z][a-z]*) \?`
	wants := make(map[string]string)
	wants["galaxyUnits"] = "glob prok"
	wants["commodity"] = "Silver"

	DoTest(input, regex, wants, t)
}
func DoTest(input, regex string, wants map[string]string, t *testing.T) {

	got := MatchInput(input, regex)

	for key := range got {
		if wants[key] != got[key] {
			t.Errorf("MatchInput(%s, %s) False output,\ngot param %s = %s, wants = %s", input, regex, key, got[key], wants[key])
		}
	}
}
