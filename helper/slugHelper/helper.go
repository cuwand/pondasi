package slugHelper

import "strings"

func ToSlug(value string) string {
	cleanOfSpaceString := strings.ToLower(strings.TrimSpace(value))

	return strings.ReplaceAll(cleanOfSpaceString, " ", "-")
}
