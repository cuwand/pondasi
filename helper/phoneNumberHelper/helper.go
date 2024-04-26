package phoneNumberHelper

import (
	"fmt"
	"regexp"
)

var cleanRegex *regexp.Regexp

func init() {
	cleanRegexCompile, err := regexp.Compile("[^0-9a-zA-Z]+")

	if err != nil {
		panic(err)
	}

	cleanRegex = cleanRegexCompile
}

func ReformatNumber(phoneNumber string) string {
	phoneNumber = cleanRegex.ReplaceAllString(phoneNumber, "")

	if len(phoneNumber) < 5 {
		panic(fmt.Sprintf("Lenght Number not Valid, Length Number: %v", len(phoneNumber)))
	}

	if phoneNumber[0:3] == "620" {
		phoneNumber = "0" + phoneNumber[3:]
	} else if phoneNumber[0:2] == "62" {
		phoneNumber = "0" + phoneNumber[2:]
	} else if phoneNumber[0:1] != "0" {
		phoneNumber = "0" + phoneNumber
	}

	return phoneNumber
}

func ReformatNumberID(phoneNumber string) string {
	if len(phoneNumber) < 5 {
		panic(fmt.Sprintf("Lenght Number not Valid, Length Number: %v", len(phoneNumber)))
	}

	return phoneNumber
}
