package stringHelper

func IsPresent(str *string) bool {
	if str != nil && len(*str) > 0 {
		return true
	}

	return false
}
