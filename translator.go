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
	var decimal float32

	for _, r := range roman {
		switch r {
		case 'I':
			decimal += 1
		case 'V':
			decimal += 5
		case 'X':
			decimal += 10
		case 'L':
			decimal += 50
		case 'C':
			decimal += 100
		case 'D':
			decimal += 500
		case 'M':
			decimal += 1000
		}
	}

	return decimal
}

//Doing the commands
func (s *System) doAction(input, action, regex string) {
	switch action {
	case "setGalaxyUnit":
		groups := MatchInput(input, regex)
		//fmt.Println(groups)
		var galaxyUnit = groups["galaxyUnit"]
		var romanNum = groups["romanNum"]

		//defining the galaxy unit
		s.galaxyUnit[galaxyUnit] = romanNum
		//fmt.Println(s.galaxyUnit)

	case "setGalaxyUnitCredit":

		groups := MatchInput(input, regex)
		//fmt.Println(groups)
		var mineral = groups["mineral"]

		galaxyUnits := strings.Split(groups["galaxyUnits"], " ")
		println(galaxyUnits, groups["galaxyUnits"], "sas")
		var romanStr string
		for i := range galaxyUnits {
			romanStr += s.galaxyUnit[galaxyUnits[i]]
		}

		//convert roman to decimal
		romanNum := RomanToDecimal(romanStr)
		fmt.Println(romanNum, romanStr, "romanNum")

		credit, err := strconv.Atoi(groups["credit"])
		if err != nil {
			panic(err)
		}
		//defining the galaxy unit
		s.galaxyCredit[mineral] = float32(credit) / romanNum
		fmt.Println(s.galaxyCredit, "zzz")
	case "getGalaxyUnitCredit":
	case "getGalaxyCredit":
	}
}
