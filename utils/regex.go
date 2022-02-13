package utils

import "regexp"

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
