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
	acronymLettersFound := make([]bool, len(acronymChars))
	productsHavingLetter := make([]bool, len(productName))

	fmt.Println(acronym)
	fmt.Println(acronymChars)

	hitsOrder := make([][]int, len(productName))
	for i := range hitsOrder {
		hitsOrder[i] = make([]int, 0)
	}

	productNameLettersMatched := make(map[int][]bool)
	for i, word := range productName {
		productNameLettersMatched[i] = make([]bool, len(word))
	}

	isProductHitInOrder := make([]bool, len(productName))

	fmt.Printf("Testing %v\n", acronymChars)
	fmt.Printf("Against %v\n", productName)
	fmt.Println(productNameLettersMatched)

	for wordIndex, productNameWord := range productName {
		// Checking word GISi, Zombie, Tracker, 3000
		fmt.Println("....")

		// for letterIndex, acronymLetter := range acronymChars {

		for letterIndex := 0; letterIndex < len(acronymChars); letterIndex++ {
			acronymLetter := acronymChars[letterIndex]
			// Checking letters G, I, Z, T, K, 3 in word GISi, Zombie, Tracker, 3000
			fmt.Println(acronymLetter + " " + strconv.FormatBool(acronymLettersFound[letterIndex]))
			if strings.Contains(productNameWord, acronymLetter) { // Check if the word contains the acronym letter
				if !acronymLettersFound[letterIndex] { // Check if the acronym letter has been found

					fmt.Printf("acronym[%v] : %v is contained in productNameWord[%v][%v] : [%v][%v]\n", letterIndex, acronymLetter, wordIndex, strings.Index(productNameWord, acronymLetter), productNameWord, string(productNameWord[strings.Index(productNameWord, acronymLetter)]))

					acronymLettersFound[letterIndex] = true
					productsHavingLetter[wordIndex] = true
					hitsOrder[wordIndex] = append(hitsOrder[wordIndex], strings.Index(productNameWord, acronymLetter))
					productNameLettersMatched[wordIndex][strings.Index(productNameWord, acronymLetter)] = true
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
	for _, matches := range productNameLettersMatched {
		fmt.Println(matches)
	}

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

	return true
}
