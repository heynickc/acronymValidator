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
	// squashedProductNameChars := strings.Split(squashedProductName, "")
	// Container for valid acronym character matches against productName
	productNamesRepresented := make([]bool, len(productName))

	validAcronymChars := make([]string, 0)

	acronymChars := strings.Split(strings.ToLower(acronym), "")

	for _, acronymChar := range acronymChars {

		for productIndex, product := range productName {
			fmt.Printf("Testing acronymChar %v against product %v\n", acronymChar, product)
			if strings.Contains(product, acronymChar) {
				fmt.Printf("%v is located in %v\n", acronymChar, product)
				productNamesRepresented[productIndex] = true
				break
			}
		}
		if strings.Contains(squashedProductName, acronymChar) {

			// acronymCharLocation := strings.Index(squashedProductName, acronymChar)
			// Store valid acronym letter
			validAcronymChars = append(validAcronymChars, acronymChar)
			// Truncate product name up to found acronym character using slice syntax
			// squashedProductNameChars = squashedProductNameChars[acronymCharLocation:]
			// squashedProductName = strings.Join(squashedProductNameChars, "")

			fmt.Println(squashedProductName)
			fmt.Println(validAcronymChars)
			continue
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
