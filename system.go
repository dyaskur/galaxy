package main

import (
	"fmt"
	"github.com/dyaskur/galaxy/utils"
	"regexp"
	"strconv"
)

type System struct {
	galaxyUnit   map[string]string
	galaxyCredit map[string]float64
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
	s.galaxyCredit = make(map[string]float64)
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

//DoAction Doing the commands
func (s *System) DoAction(input, action, regex string) {
	groups := utils.MatchInput(input, regex)
	switch action {
	case "setGalaxyUnit":
		var galaxyUnit = groups["galaxyUnit"]
		var romanNum = groups["romanNum"]

		//defining the galaxy unit
		s.galaxyUnit[galaxyUnit] = romanNum

	case "setGalaxyUnitCredit":
		var commodity = groups["commodity"]

		romanNum := utils.UnitsToDecimal(groups["galaxyUnits"], s.galaxyUnit)

		credit, err := strconv.ParseFloat(groups["credit"], 64)
		if err != nil {
			panic(err)
		}
		//defining the galaxy unit
		s.galaxyCredit[commodity] = credit / romanNum
		//answer how much question (e.g. how much is pish tegj glob glob ? )
	case "getGalaxyUnitCredit":
		romanNum := utils.UnitsToDecimal(groups["galaxyUnits"], s.galaxyUnit)

		//Add output
		if romanNum == -1 {
			s.AddOutput("I have no idea what you are talking about")
		} else {
			s.AddOutput(fmt.Sprintf("%s is %v", groups["galaxyUnits"], romanNum))
		}

	case "getGalaxyCredit": //answer how many Credits question (e.g. how many Credits is glob prok Silver ? )
		romanNum := utils.UnitsToDecimal(groups["galaxyUnits"], s.galaxyUnit)

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
