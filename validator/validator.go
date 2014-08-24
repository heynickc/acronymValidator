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
	for len(acronymCharsCopy) > 0 {
		var char string
		char, acronymCharsCopy = Pop(acronymCharsCopy)
		for i := 0; i < len(productName); i++ {
			if strings.Contains(productName[i], char) {
				// fmt.Println(char)
				// fmt.Println(acronymCharsCopy)

				err := nameBucketList.AddItemAtIndex(i, char)
				if err != nil {
					fmt.Println(err)
				} else {
					productName[i] = productName[i][strings.Index(productName[i], char):]
					// fmt.Println(productName)
					break
				}
			}
		}
	}

	fmt.Println(nameBucketList)

	return true
}

func Pop(strings []string) (string, []string) {
	popped, strings := strings[0], strings[1:len(strings)]
	return popped, strings
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}
