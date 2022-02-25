package slugHelper

import "strings"

func Generate(value string) string {
	cleanOfSpaceString := strings.ToLower(strings.TrimSpace(value))

	return strings.ReplaceAll(cleanOfSpaceString, " ", "-")
}