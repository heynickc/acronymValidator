package validator

import (
	"fmt"
	ds "github.com/heynickc/acronym_validator/data_structures"
	utils "github.com/heynickc/acronym_validator/utilities"
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
		char, acronymChars = utils.Pop(acronymChars)
		for i, name := range productName {
			if strings.Contains(name, char) {
				err := nameBucketList.AddItemAtIndex(i, char)
				if err != nil {
					continue
				} else {
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
