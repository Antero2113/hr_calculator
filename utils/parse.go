package utils

import "strings"

func ParseDepartment(input string) (string, string) {
	input = strings.TrimSpace(input)
	directions := []string{"СЕВЕР", "ЮГ", "ЗАПАД", "ВОСТОК"}

	upper := strings.ToUpper(input)

	for _, dir := range directions {
		if strings.HasSuffix(upper, dir) {
			dept := strings.TrimSpace(input[:len(input)-len(dir)])
			return dept, dir
		}
	}

	return input, ""
}