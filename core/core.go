package core

import "regexp"

func findIPAddress(input string) string {
	numBlock := "(((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]?[0-9])[.]){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]?[0-9]))"

	regEx := regexp.MustCompile(numBlock)
	return regEx.FindString(input)
}

func findMACAddress(input string) string {
	regexPattern := "(([A-Fa-f0-9]{2}[:]){5}([A-Fa-f0-9]{2}))"

	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}
