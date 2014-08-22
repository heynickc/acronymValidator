package main

import (
	"./validator"
	"fmt"
	"strings"
)

// Requirements:
// Each product name has at least one abbreviation in the acronym
// Each abbreviation in the acronym has a respective product name
// The letters in the acronym appear in order for the entire string of product names

func main() {

	productName := strings.Fields(strings.ToLower("GISi Zombie Tracker 3000"))
	validAcronym := "GIZTK3"
	// invalidAcronym := "GZT3k"

	fmt.Println(validator.IsValid(validAcronym, productName))
	// fmt.Println(validator.IsValid(invalidAcronym, productName))
}
