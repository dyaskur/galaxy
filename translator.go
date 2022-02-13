package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type System struct {
	galaxyUnit   map[string]string
	galaxyCredit map[string]float32
	//List of variable in regex
	regexVars map[string]string
	//List of regex formulas/patterns
	regexFormulas []RegexFormula
	//list of output that will be printed
	output []string
}

//RegexFormula for save all regex and command for queries.
type RegexFormula struct {
	formula string
	action  string
}

//Init variable
func (s *System) Init() {
	s.galaxyUnit = make(map[string]string)
	s.galaxyCredit = make(map[string]float32)
	s.regexVars = make(map[string]string)

	s.regexVars["galaxyUnit"] = `(?P<galaxyUnit>([a-z]+))`
	s.regexVars["galaxyUnits"] = `(?P<galaxyUnits>([a-z ]+))`
	s.regexVars["roman"] = `(?P<romanNum>([IVXLCDM]+))`
	s.regexVars["commodity"] = `(?P<commodity>([A-Z][a-z]*))`
	s.regexVars["credit"] = `(?P<credit>([0-9]+))`

	s.regexFormulas = []RegexFormula{
		{
			formula: s.regexVars["galaxyUnit"] + ` is ` + s.regexVars["roman"] + ``,
			action:  "setGalaxyUnit",
		},
		{
			formula: s.regexVars["galaxyUnits"] + ` ` + s.regexVars["commodity"] + ` is ` + s.regexVars["credit"] + ` Credits$`,
			action:  "setGalaxyUnitCredit",
		},
		{
			formula: `^how much is ` + s.regexVars["galaxyUnits"] + ` \?$`,
			action:  "getGalaxyUnitCredit",
		},
		{
			formula: `^how many Credits is ` + s.regexVars["galaxyUnits"] + ` ` + s.regexVars["commodity"] + ` \?$`,
			action:  "getGalaxyCredit",
		},
		{
			formula: `^how much is ` + s.regexVars["galaxyUnits"] + `\?$`,
			action:  "getGalaxyUnitCredit",
		},
		{
			formula: `^how many Credits is ` + s.regexVars["galaxyUnits"] + ` ` + s.regexVars["commodity"] + `\?$`,
			action:  "getGalaxyCredit",
		},
	}
}

//Translate or convert input to command/action
func (s *System) Translate(input string) {
	//flag for input is understandable by regex formula or not
	isUnderstandable := false

	for _, reg := range s.regexFormulas {
		re, err := regexp.Compile(reg.formula)
		if err != nil {
			panic(err)
		}

		if re.Match([]byte(input)) {
			s.DoAction(input, reg.action, reg.formula)
			isUnderstandable = true //this input is understandable by regex formula
			break
		}
	}

	//If the system don't understand the input, print output
	if isUnderstandable == false {
		s.AddOutput("I have no idea what you are talking about")
	}
}

// MatchInput match the input with regex and get value of the variable/data
func MatchInput(input, regex string) map[string]string {
	re := regexp.MustCompile(regex)

	match := re.FindStringSubmatch(input)

	groups := make(map[string]string)

	for i, groupName := range re.SubexpNames() {
		if i == 0 {
			continue
		}

		groups[groupName] = match[i]
	}

	return groups
}

// RomanToDecimal  to convert Roman to decimal
func RomanToDecimal(roman string) float32 {
	var romanValue float32
	var result float32

	romanValues := make(map[int32]float32)
	romanValues['I'] = 1
	romanValues['V'] = 5
	romanValues['X'] = 10
	romanValues['L'] = 50
	romanValues['C'] = 100
	romanValues['D'] = 500
	romanValues['M'] = 1000

	leadingOnes := make(map[int32]bool)
	leadingOnes['I'] = true
	leadingOnes['V'] = false
	leadingOnes['X'] = true
	leadingOnes['L'] = false
	leadingOnes['C'] = true
	leadingOnes['D'] = false
	leadingOnes['M'] = true

	//Temporary variables for check duplicate
	var countDuplicate int
	var duplicateChar int32
	for i, r := range roman {
		//The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.) "D", "L", and "V" can never be repeated.
		//Check duplicate
		if duplicateChar != r {
			duplicateChar = r
			countDuplicate = 1
		} else {
			if leadingOnes[r] == true {
				countDuplicate++
			} else {
				result = -1
				return result
			}
		}

		if countDuplicate > 3 {
			result = -1
			return result
		}
		//end of check duplicate
		romanValue = romanValues[r]

		//"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.
		//Only one small-value symbol may be subtracted from any large-value symbol.
		if leadingOnes[r] && i != len(roman)-1 && romanValue < romanValues[int32(roman[i+1])] {
			result -= romanValue
		} else {
			result += romanValue
		}
	}

	return result
}

//DoAction Doing the commands
func (s *System) DoAction(input, action, regex string) {
	groups := MatchInput(input, regex)
	switch action {
	case "setGalaxyUnit":
		var galaxyUnit = groups["galaxyUnit"]
		var romanNum = groups["romanNum"]

		//defining the galaxy unit
		s.galaxyUnit[galaxyUnit] = romanNum

	case "setGalaxyUnitCredit":
		var commodity = groups["commodity"]

		galaxyUnits := strings.Split(groups["galaxyUnits"], " ")
		var romanStr string
		for i := range galaxyUnits {
			romanStr += s.galaxyUnit[galaxyUnits[i]]
		}

		//convert roman to decimal
		romanNum := RomanToDecimal(romanStr)

		credit, err := strconv.Atoi(groups["credit"])
		if err != nil {
			panic(err)
		}
		//defining the galaxy unit
		s.galaxyCredit[commodity] = float32(credit) / romanNum
		//answer how much question (e.g. how much is pish tegj glob glob ? )
	case "getGalaxyUnitCredit":
		var galaxyUnits = strings.Split(groups["galaxyUnits"], " ")
		var romanStr string
		for i := range galaxyUnits {
			romanStr += s.galaxyUnit[galaxyUnits[i]]
		}

		//convert roman to decimal
		romanNum := RomanToDecimal(romanStr)

		//Add output
		if romanNum == -1 {
			s.AddOutput("I have no idea what you are talking about")
		} else {
			s.AddOutput(fmt.Sprintf("%s is %v", groups["galaxyUnits"], romanNum))
		}

	case "getGalaxyCredit": //answer how many Credits question (e.g. how many Credits is glob prok Silver ? )
		var galaxyUnits = strings.Split(groups["galaxyUnits"], " ")
		var romanStr string
		for i := range galaxyUnits {
			romanStr += s.galaxyUnit[galaxyUnits[i]]
		}

		//convert roman to decimal
		romanNum := RomanToDecimal(romanStr)

		//if the galaxy units is not valid, e.g too many duplicate unit
		if romanNum == -1 {
			s.AddOutput("Your units are not valid")
		} else {
			var commodity = groups["commodity"]
			//Add output
			totalCredit := s.galaxyCredit[commodity] * romanNum
			output := fmt.Sprintf("%s %s is %v Credits", groups["galaxyUnits"], commodity, totalCredit)
			s.AddOutput(output)
		}
	}
}

//AddOutput Function for add output
func (s *System) AddOutput(output string) {
	s.output = append(s.output, output)
}
