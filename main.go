package main

import (
	"./validator"
	"fmt"
	"strings"
)

func main() {

	productName := strings.Fields(strings.ToLower("GISi Zombie Tracker 3000"))
	validAcronym := "GIZTK3"
	// invalidAcronym := "GZT3k"

	fmt.Println(validator.IsValid(validAcronym, productName))
	// fmt.Println(validator.IsValid(invalidAcronym, productName))
}
