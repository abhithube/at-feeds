package parser

import "regexp"

func GetNamedGroups(regex *regexp.Regexp, s string) map[string]string {
	match := regex.FindStringSubmatch(s)
	if match == nil {
		return nil
	}

	result := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
