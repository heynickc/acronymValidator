package main

import (
	// "./validator"
	"fmt"
	"sort"
	"strings"
)

func main() {

	productName := strings.Fields(strings.ToLower("GISi Zombie Tracker 3000"))
	validAcronym := strings.Split(strings.ToLower("GIZTK3"), "")
	// invalidAcronym := "GZT3k"

	acronymLettersContained := make([]bool, len(validAcronym))
	productsHavingLetter := make([]bool, len(productName))
	hitsOrder := make([][]int, len(productName))
	for i := range hitsOrder {
		hitsOrder[i] = make([]int, 0)
	}
	isProductHitInOrder := make([]bool, len(productName))

	fmt.Println(acronymLettersContained)
	fmt.Println(productsHavingLetter)
	fmt.Println(isProductHitInOrder)

	for wordIndex, productNameWord := range productName {
		// Checking word GISi, Zombie, Tracker, 3000
		for letterIndex, acronymLetter := range validAcronym {
			// Checking letters G, I, Z, T, K, 3 in word GISi
			if strings.Contains(productNameWord, acronymLetter) { // Check if the word contains the acronym letter
				if !acronymLettersContained[letterIndex] { // Check if the acronym letter has been found
					fmt.Printf("acronym[%v] : %v is contained in productNameWord[%v][%v] : [%v][%v]\n", letterIndex, acronymLetter, wordIndex, strings.Index(productNameWord, acronymLetter), productNameWord, string(productNameWord[strings.Index(productNameWord, acronymLetter)]))
					acronymLettersContained[letterIndex] = true
					productsHavingLetter[wordIndex] = true
					hitsOrder[wordIndex] = append(hitsOrder[wordIndex], strings.Index(productNameWord, acronymLetter))
				}
			}
		}
	}

	for hitsIndex, hits := range hitsOrder {
		if sort.IntsAreSorted(hits) {
			isProductHitInOrder[hitsIndex] = true
		}
	}

	fmt.Println(acronymLettersContained)
	fmt.Println(productsHavingLetter)
	fmt.Println(isProductHitInOrder)
}
