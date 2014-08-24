package validator

import (
	"fmt"
	// "math"
	// "sort"
	ds "github.com/heynickc/acronym_validator/data_structures"
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
	// productNamesRepresentedFlag := make([]bool, len(productName))
	// productNamesRepresented := make([]string, 0)

	// For checking the order of acronym character matches against productName
	squashedProductName := strings.Join(productName, "")
	squashedProductNameChars := strings.Split(squashedProductName, "")
	validAcronymChars := make([]string, 0)

	// for i := 0; i < len(productName); i++ {
	// 	acronymCharsCopy := make([]string, len(acronymChars))
	// 	copy(acronymCharsCopy, acronymChars)
	// 	container := make([]string, 0)
	// 	for len(acronymCharsCopy) > 0 {
	// 		var char string
	// 		char, acronymCharsCopy = Pop(acronymCharsCopy)
	// 		if strings.Contains(productName[i], char) {
	// 			// Insert at beginning
	// 			container = InsertStringSlice(container, []string{char}, 0)
	// 		}
	// 	}
	// 	fmt.Println(container)
	// }

	var nameBucketList = ds.BucketList{make([]ds.Bucket, 0)}

	for _, name := range productName {
		nameBucketList.AddBucket(ds.Bucket{name, make([]string, 0), 0, 0})
	}

	if len(acronymChars) >= len(productName) {
		nameBucketList.SetCapacities(len(acronymChars))
	} else {
		return false
	}

	acronymCharsCopy := make([]string, len(acronymChars))
	copy(acronymCharsCopy, acronymChars)
	fmt.Println(acronymCharsCopy)
	for i := 0; i < len(productName); i++ {
		for len(acronymCharsCopy) > 0 {
			var char string
			char, acronymCharsCopy = Pop(acronymCharsCopy)
			if strings.Contains(productName[i], char) {
				nameBucketList.AddItemAtIndex(i, char)
			}
		}
	}

	fmt.Println(nameBucketList)

	for _, acronymChar := range acronymChars {
		if strings.Contains(squashedProductName, acronymChar) {
			acronymCharLocation := strings.Index(squashedProductName, acronymChar)
			validAcronymChars = append(validAcronymChars, acronymChar)
			squashedProductNameChars = squashedProductNameChars[acronymCharLocation:]
			squashedProductName = strings.Join(squashedProductNameChars, "")
		}
	}

	// for _, isRepresented := range productNamesRepresentedFlag {
	// 	if !isRepresented {
	// 		return false
	// 	}
	// }

	// // Checks that there is an acronym letter for every productName string
	// // fmt.Printf("Original product names: %v, Order of matched product names: %v\n", productName, productNamesRepresented)
	// if !strings.EqualFold(strings.Join(productName, ""), strings.Join(productNamesRepresented, "")) {
	// 	return false
	// }

	// Makes sure the overall acronym is in order, without considering productName boundaries
	// fmt.Printf("Original acronym: %v, Valid acronym characters: %v\n", strings.ToLower(acronym), strings.Join(validAcronymChars, ""))
	if !strings.EqualFold(acronym, strings.Join(validAcronymChars, "")) { //EqualFold compares lower-case equality
		return false
	}

	return true
}

func Pop(strings []string) (string, []string) {
	popped, strings := strings[len(strings)-1], strings[:len(strings)-1]
	return popped, strings
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}
