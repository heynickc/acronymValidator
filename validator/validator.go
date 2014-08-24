package validator

import (
	"fmt"
	ds "github.com/heynickc/acronyValidator/data_structures"
	utils "github.com/heynickc/acronymValidator/utilities"
	"strings"
)

func IsValid(acronym string, productName []string) bool {
	fmt.Printf("Testing %v against %v\n", acronym, productName)

	var acronymChars []string
	productName, acronymChars = ProcessInputs(acronym, productName)

	var nameBucketList = GetProductNameBuckets(productName)

	if len(acronymChars) >= len(productName) {
		nameBucketList.SetCapacities(len(acronymChars))
	} else {
		return false
	}

	for len(acronymChars) > 0 {
		var char string
		// Take each letter from acronym and test it against the product name
		char, acronymChars = utils.Pop(acronymChars)
		for i, name := range productName {
			// Product name has the letter, try to insert it into the name's bucket
			if strings.Contains(name, char) {
				// If the bucket needs more room, it will resize
				// If the bucket can't resize, it will return err
				err := nameBucketList.AddItemAtIndex(i, char)
				if err != nil {
					// Bucket can't resize, so skip to the next word
					continue
				} else {
					// Keeps things in order by truncating before the match
					productName[i] = productName[i][strings.Index(name, char)+1:]
					break
				}
			} else {
				productName[i] = ""
				continue
			}
		}
	}

	matchedAcronymChars := nameBucketList.GetAllItemsSquashed()
	// Checks the buckets against the input acronym
	// Returns false to IsValid if the conditions aren't met
	if !strings.EqualFold(acronym, matchedAcronymChars) || !nameBucketList.AllBucketsHaveItems() {
		return false
	}

	return true
}

func ProcessInputs(acronym string, productName []string) ([]string, []string) {
	for i, _ := range productName {
		productName[i] = strings.ToLower(productName[i])
	}
	acronymChars := strings.Split(strings.ToLower(acronym), "")
	return productName, acronymChars
}

func GetProductNameBuckets(productName []string) ds.BucketList {
	var nameBucketList = ds.BucketList{make([]ds.Bucket, 0)}
	for _, name := range productName {
		nameBucketList.AddBucket(ds.Bucket{name, make([]string, 0), 0, 0})
	}
	return nameBucketList
}
