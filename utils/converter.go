package utils

import "strings"

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

func UnitsToDecimal(units string, definedUnits map[string]string) float32 {
	galaxyUnits := strings.Split(units, " ")
	var romanStr string
	for i := range galaxyUnits {
		romanStr += definedUnits[galaxyUnits[i]]
	}

	//convert roman to decimal
	romanNum := RomanToDecimal(romanStr)
	return romanNum
}
