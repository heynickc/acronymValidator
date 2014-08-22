package validator

import (
	"fmt"
	// "sort"
	// "strconv"
	"strings"
)

func IsValid(acronym string, productName []string) bool {
	fmt.Println("Testing %v against %v", acronym, productName)

	for i, _ := range productName {
		productName[i] = strings.ToLower(productName[i])
	}

	squashedProductName := strings.Join(productName, "")
	squashedProductNameSplit := strings.Split(squashedProductName, "")
	fmt.Println(squashedProductName)

	for wordIndex, productNameWord := range productName {
		// Checking word GISi, Zombie, Tracker, 3000
		fmt.Println("Testing - " + productNameWord)
		for letterIndex, acronymLetter := range acronymChars {
			// Checking letters G, I, Z, T, K, 3 in word GISi, Zombie, Tracker, 3000

			fmt.Printf("\"%v\"\n", letterIndex)

			if strings.ContainsAny(productNameWord, acronymChars[letterIndex]) { // Check if the word contains any of the acronym letters

				productNameLocation := strings.Index(productNameWord, acronymLetter)
				productNameSplit := strings.Split(productNameWord, "")
				productNameSplit = RemoveStringSliceCopy(productNameSplit, productNameLocation, productNameLocation+1)
				productNameWord = strings.Join(productNameSplit, "")

				letterLocation := strings.Index(squashedProductName, acronymLetter)
				validAcronymChars = append(validAcronymChars, squashedProductNameSplit[letterLocation])
				squashedProductNameSplit = RemoveStringSliceCopy(squashedProductNameSplit, letterLocation, letterLocation+1)
				squashedProductName = strings.Join(squashedProductNameSplit, "")

			}
		}
	}

	fmt.Printf("%v == %v?\n", strings.Join(strings.Split(strings.ToLower(acronym), ""), ""), strings.Join(validAcronymChars, ""))
	if !strings.EqualFold(strings.Join(strings.Split(strings.ToLower(acronym), ""), ""), strings.Join(validAcronymChars, "")) {
		return false
	}

	return true
}

func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}
