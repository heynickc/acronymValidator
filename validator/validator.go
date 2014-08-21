package validator

import (
	"strings"
)

func IsValid(acronym string, productName []string) bool {
	if IsValidCharCount(acronym, productName) && IsValidContainment(acronym, productName) && IsValidCharOrder(acronym, productName) {
		return true
	}
	return false
}

func IsValidCharCount(acronym string, productName []string) bool {
	acronymSplit := strings.Split(acronym, "")
	if len(acronymSplit) < len(productName) {
		return false
	}
	return true
}

func IsValidContainment(acronym string, productName []string) bool {
	containsAcronym := make([]bool, len(productName))
	for i, s := range productName {
		containsAcronym[i] = strings.ContainsAny(strings.ToLower(s), strings.ToLower(acronym))
	}
	for _, b := range containsAcronym {
		if !b {
			return false
		}
	}
	return true
}

func IsValidCharOrder(acronym string, productName []string) bool {
	squashedProductName := Squash(productName)
	acronymSplit := strings.Split(strings.ToLower(acronym), "")
	indexCounter := -1
	for _, char := range acronymSplit {
		if indexCounter <= strings.Index(squashedProductName, char) {
			indexCounter = strings.Index(squashedProductName, char)
		} else {
			return false
		}
	}
	return true
}

func Squash(productName []string) string {
	return strings.ToLower(strings.Join(productName, ""))
}
