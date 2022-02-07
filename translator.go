package main

import (
	"fmt"
	"regexp"
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
	s.regexVars["galaxyUnits"] = `(?P<galaxyUnit>([a-z ]+))`
	s.regexVars["roman"] = `(?P<romanNum>([IVXLCDM]+))`
	s.regexVars["mineral"] = `(?P<mineral>([A-Z][a-z]*))`
	s.regexVars["credit"] = `(?P<credit>([0-9]+))`

	s.regexFormulas = []RegexFormula{
		{
			formula: s.regexVars["galaxyUnit"] + ` is ` + s.regexVars["roman"] + ``,
			action:  "setGalaxyUnit",
		},
		{
			formula: s.regexVars["galaxyUnits"] + ` ` + s.regexVars["mineral"] + ` is ` + s.regexVars["credit"] + ` Credits$`,
			action:  "setGalaxyUnitCredit",
		},
		{
			formula: `^how much is ` + s.regexVars["galaxyUnits"] + ` \?$`,
			action:  "getGalaxyUnitCredit",
		},
		{
			formula: `^how many Credits is ` + s.regexVars["galaxyUnits"] + ` ` + s.regexVars["mineral"] + ` \?$`,
			action:  "getGalaxyCredit",
		},
	}
}

func (s *System) translator(input string) string {
	var output string

	r := regexp.MustCompile(s.regexVars["galaxyUnit"] + ` is ` + s.regexVars["roman"])

	// Using FindStringSubmatch you are able to access the
	// individual capturing groups
	fmt.Printf("%#v\n", r.FindStringSubmatch(input))
	fmt.Printf("%#v\n", r.SubexpNames())

	return output
}

//Translate input to command/action
func (s *System) Translate(input string) {
	//flag for query is already process or not
	isUnderstandable := false

	for _, reg := range s.regexFormulas {
		re, err := regexp.Compile(reg.formula)
		if err != nil {
			panic(err)
		}

		fmt.Println(re.FindStringSubmatch(input))

		if re.Match([]byte(input)) {
			s.doAction(input, reg.action, reg.formula)
			isUnderstandable = true
			break
		}
	}

	//If the system don't understand the input, print output
	if isUnderstandable == false {
		println("I have no idea what you are talking about")
	}
}

//Doing the commands
func (s *System) doAction(line, action, regex string) {
	switch action {
	case "setGalaxyUnit":
		re := regexp.MustCompile(regex)

		match := re.FindStringSubmatch(line)

		groups := make(map[string]string)

		for i, groupName := range re.SubexpNames() {
			if i == 0 {
				continue
			}

			groups[groupName] = match[i]
		}
		//fmt.Println(groups)
		var galaxyUnit = groups["galaxyUnit"]
		var romanNum = groups["romanNum"]

		//defining the galaxy unit
		s.galaxyUnit[galaxyUnit] = romanNum
		fmt.Println(s.galaxyUnit)

	case "setGalaxyUnitCredit":
	case "getGalaxyUnitCredit":
	case "getGalaxyCredit":
	}
}
