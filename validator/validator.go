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

	acronymChars := strings.Split(strings.ToLower(acronym), "")
	// Tests if all the letters are used in the productName
	acronymLettersFound := make([]bool, len(acronymChars))
	// Tests if each productName contains a letter
	productsHavingLetter := make([]bool, len(productName))

	validAcronymChars := make([]string, 0)

	for wordIndex, productNameWord := range productName {
		// Checking word GISi, Zombie, Tracker, 3000
		fmt.Println("Testing - " + productNameWord)
		for letterIndex, acronymLetter := range acronymChars {
			// for i := 0; i < len(acronymChars); i++ {
			// Checking letters G, I, Z, T, K, 3 in word GISi, Zombie, Tracker, 3000
			if strings.ContainsAny(productNameWord, acronymChars[letterIndex]) { // Check if the word contains any of the acronym letters
				fmt.Printf("acronym[%v] : %v is contained in productNameWord[%v][%v] : [%v][%v]\n", letterIndex, acronymChars[letterIndex], wordIndex, strings.Index(productNameWord, acronymLetter), productNameWord, string(productNameWord[strings.Index(productNameWord, acronymChars[letterIndex])]))
				fmt.Println(acronymChars[letterIndex])
				acronymLettersFound[letterIndex] = true
				productsHavingLetter[wordIndex] = true

				acronymChars[letterIndex] = ""

				productNameLocation := strings.Index(productNameWord, acronymLetter)
				productNameSplit := strings.Split(productNameWord, "")
				productNameSplit = RemoveStringSliceCopy(productNameSplit, productNameLocation, productNameLocation+1)
				productNameWord = strings.Join(productNameSplit, "")

				letterLocation := strings.Index(squashedProductName, acronymLetter)

				validAcronymChars = append(validAcronymChars, squashedProductNameSplit[letterLocation])
				squashedProductNameSplit = RemoveStringSliceCopy(squashedProductNameSplit, letterLocation, letterLocation+1)
				squashedProductName = strings.Join(squashedProductNameSplit, "")

				// fmt.Println(squashedProductName)
				// fmt.Println(validAcronymChars)
				// fmt.Println(acronymChars)
				// fmt.Println(productNameWord)
			}
		}
	}

	fmt.Println(acronymLettersFound)
	for _, a := range acronymLettersFound {
		if !a {
			return false
		}
	}

	fmt.Println(productsHavingLetter)
	for _, b := range productsHavingLetter {
		if !b {
			return false
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
