package validator

import (
	"fmt"
	// "sort"
	// "strconv"
	"strings"
)

func IsValid(acronym string, productName []string) bool {
	fmt.Printf("Testing %v against %v\n", acronym, productName)

	// Process the inputs
	for i, _ := range productName {
		productName[i] = strings.ToLower(productName[i])
	}
	acronymChars := strings.Split(strings.ToLower(acronym), "")

	// Container for valid acronym character matches against productName
	productNamesRepresentedFlag := make([]bool, len(productName))
	productNamesRepresented := make([]string, 0)

	// For checking the order of acronym character matches against productName
	squashedProductName := strings.Join(productName, "|")
	squashedProductNameChars := strings.Split(squashedProductName, "")
	validAcronymChars := make([]string, 0)

	for _, acronymChar := range acronymChars {
		for i := 0; i < len(productName); i++ {

			fmt.Printf("Testing acronymChar %v against product %v\n", acronymChar, productName[i])

			if product := productName[i]; strings.Contains(product, acronymChar) && !productNamesRepresentedFlag[i] {
				fmt.Printf("%v is located in %v\n", acronymChar, product)

				productNamesRepresentedFlag[i] = true
				productNamesRepresented = append(productNamesRepresented, product)
				break
			}
		}
	}

	for _, acronymChar := range acronymChars {

		if strings.Contains(squashedProductName, acronymChar) {
			acronymCharLocation := strings.Index(squashedProductName, acronymChar)
			// Store valid acronym letter
			validAcronymChars = append(validAcronymChars, acronymChar)
			// Truncate product name up to found acronym character using slice syntax
			// This keeps us from finding matches we already searched through
			squashedProductNameChars = squashedProductNameChars[acronymCharLocation:]
			squashedProductName = strings.Join(squashedProductNameChars, "")

			continue
		}
	}

	// Checks that there is an acronym letter for every productName string
	fmt.Printf("Original product names: %v, Order of matched product names: %v\n", productName, productNamesRepresented)
	if !strings.EqualFold(strings.Join(productName, ""), strings.Join(productNamesRepresented, "")) {
		return false
	}

	// Makes sure the overall acronym is in order, without considering productName boundaries
	fmt.Printf("Original acronym: %v, Valid acronym characters: %v\n", strings.ToLower(acronym), strings.Join(validAcronymChars, ""))
	if !strings.EqualFold(acronym, strings.Join(validAcronymChars, "")) { //EqualFold compares lower-case equality
		return false
	}

	return true
}
