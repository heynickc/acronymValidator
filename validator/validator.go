package validator

import (
	"fmt"
	// "sort"
	// "strconv"
	"strings"
)

func IsValid(acronym string, productName []string) bool {
	fmt.Printf("Testing %v against %v\n", acronym, productName)

	for i, _ := range productName {
		productName[i] = strings.ToLower(productName[i])
	}

	squashedProductName := strings.Join(productName, "|")
	squashedProductNameChars := strings.Split(squashedProductName, "")
	acronymChars := strings.Split(strings.ToLower(acronym), "")
	validAcronymChars := make([]string, 0)
	productNameContains := make([]bool, len(productName))

	for _, acronymChar := range acronymChars {

		for i, product := range productName {
			if string.Contains(product, acronymChar) {
				// check for containment
			}
		}
		if strings.Contains(squashedProductName, acronymChar) {

			acronymCharLocation := strings.Index(squashedProductName, acronymChar)
			// Store valid acronym letter
			validAcronymChars = append(validAcronymChars, acronymChar)
			// Remove acronym letter from the product name
			squashedProductNameChars = squashedProductNameChars[acronymCharLocation:]
			squashedProductName = strings.Join(squashedProductNameChars, "")
			// Set original acronym letter to blank
			// acronymChars[i] = ""

			fmt.Println(squashedProductName)
			fmt.Println(validAcronymChars)
		}
	}

	// fmt.Printf("%v == %v?\n", strings.Join(strings.Split(strings.ToLower(acronym), ""), ""), strings.Join(validAcronymChars, ""))
	// if !strings.EqualFold(strings.Join(strings.Split(strings.ToLower(acronym), ""), ""), strings.Join(validAcronymChars, "")) {
	// 	return false
	// }

	return true
}

func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}
