package validator

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func IsValid(acronym string, productName []string) bool {
	for i, _ := range productName {
		productName[i] = strings.ToLower(productName[i])
	}

	acronymChars := strings.Split(strings.ToLower(acronym), "")
	// Tests if all the letters are used in the productName
	acronymLettersFound := make([]bool, len(acronymChars))
	// Tests if each productName contains a letter
	productsHavingLetter := make([]bool, len(productName))

	matchIndex := make([]int, len(acronymChars))

	// Tests if the matches are in order...
	hitsOrder := make([][]int, len(productName))
	for i := range hitsOrder {
		hitsOrder[i] = make([]int, 0)
	}

	// Tests which letters were matched
	productNameLettersMatched := make(map[int][]bool)
	for i, word := range productName {
		productNameLettersMatched[i] = make([]bool, len(word))
	}

	isProductHitInOrder := make([]bool, len(productName))

	fmt.Printf("Testing %v\n", acronymChars)
	fmt.Printf("Against %v\n", productName)
	// fmt.Println(productNameLettersMatched)
	// fmt.Println(matchIndex)

	flatIndex := 0

	for wordIndex, productNameWord := range productName {

		// Checking word GISi, Zombie, Tracker, 3000
		fmt.Println("Test ......")
		for letterIndex, acronymLetter := range acronymChars {

			// Checking letters G, I, Z, T, K, 3 in word GISi, Zombie, Tracker, 3000
			fmt.Println(acronymLetter + " " + strconv.FormatBool(acronymLettersFound[letterIndex]))

			if strings.Contains(productNameWord, acronymLetter) { // Check if the word contains the acronym letter

				if !acronymLettersFound[letterIndex] && !productNameLettersMatched[wordIndex][strings.Index(productNameWord, acronymLetter)] { // Check if the acronym letter has been found and that only the preceding letter has been found

					fmt.Printf("acronym[%v] : %v is contained in productNameWord[%v][%v] : [%v][%v]\n", letterIndex, acronymLetter, wordIndex, strings.Index(productNameWord, acronymLetter), productNameWord, string(productNameWord[strings.Index(productNameWord, acronymLetter)]))

					acronymLettersFound[letterIndex] = true
					productsHavingLetter[wordIndex] = true
					hitsOrder[wordIndex] = append(hitsOrder[wordIndex], strings.Index(productNameWord, acronymLetter))
					productNameLettersMatched[wordIndex][strings.Index(productNameWord, acronymLetter)] = true

					matchIndex[flatIndex] = letterIndex
					flatIndex++
				}
			}
		}
	}

	for hitsIndex, hits := range hitsOrder {
		if sort.IntsAreSorted(hits) {
			isProductHitInOrder[hitsIndex] = true
		}
	}

	fmt.Println(acronymLettersFound)
	fmt.Println(productsHavingLetter)
	fmt.Println(isProductHitInOrder)
	fmt.Println("\nTesting product enumaration...")
	for _, matches := range productNameLettersMatched {
		fmt.Println(matches)
	}
	// fmt.Println(matchIndex)

	for _, a := range acronymLettersFound {
		if !a {
			return false
		}
	}

	for _, b := range productsHavingLetter {
		if !b {
			return false
		}
	}

	for _, c := range isProductHitInOrder {
		if !c {
			return false
		}
	}

	if !sort.IntsAreSorted(matchIndex) {
		return false
	}

	return true
}
