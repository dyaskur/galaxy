package utils

//RegexFormula for save all regex and command for queries.
type RegexFormula struct {
	Formula string
	Action  string
}

func GetFormula() ([]RegexFormula, map[string]string, error) {
	var regexVars = make(map[string]string)

	regexVars["galaxyUnit"] = `(?P<galaxyUnit>([a-z]+))`
	regexVars["galaxyUnits"] = `(?P<galaxyUnits>([a-z ]+))`
	regexVars["roman"] = `(?P<romanNum>([IVXLCDM]+))`
	regexVars["commodity"] = `(?P<commodity>([A-Z][a-z]*))`
	regexVars["credit"] = `(?P<credit>([0-9]+))`

	regexFormulas := []RegexFormula{
		{
			Formula: regexVars["galaxyUnit"] + ` is ` + regexVars["roman"] + ``,
			Action:  "setGalaxyUnit",
		},
		{
			Formula: regexVars["galaxyUnits"] + ` ` + regexVars["commodity"] + ` is ` + regexVars["credit"] + ` Credits$`,
			Action:  "setGalaxyUnitCredit",
		},
		{
			Formula: `^how much is ` + regexVars["galaxyUnits"] + ` \?$`,
			Action:  "getGalaxyUnitCredit",
		},
		{
			Formula: `^how many Credits is ` + regexVars["galaxyUnits"] + ` ` + regexVars["commodity"] + ` \?$`,
			Action:  "getGalaxyCredit",
		},
		{
			Formula: `^how much is ` + regexVars["galaxyUnits"] + `\?$`,
			Action:  "getGalaxyUnitCredit",
		},
		{
			Formula: `^how many Credits is ` + regexVars["galaxyUnits"] + ` ` + regexVars["commodity"] + `\?$`,
			Action:  "getGalaxyCredit",
		},
	}
	return regexFormulas, regexVars, nil
}
