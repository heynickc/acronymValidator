package data_structures

import (
	"fmt"
	"testing"
)

var bucketList = BucketList{[]Bucket{
	Bucket{"GISi", make([]string, 0), 0, 0},
	Bucket{"Zombie", make([]string, 0), 0, 0},
	Bucket{"Tracker", make([]string, 0), 0, 0},
	Bucket{"3000", make([]string, 0), 0, 0},
}}

func TestSetCapacity(t *testing.T) {
	bucketList.SetCapacities(15)
	// fmt.Println(bucketList)
}

func TestGetTotalCapacity(t *testing.T) {
	// fmt.Println(bucketList.GetTotalCapacity()) // should equal 15
}

func TestGetTotalCapacityAfterIndex(t *testing.T) {
	bucketList.SetCapacities(15)
	// fmt.Println(bucketList.GetTotalCapacityAfterIndex(2)) // should equal 3 - 4th item
}

func TestSetCapacitiesAfterIndex(t *testing.T) {
	bucketList.SetCapacities(15)
	bucketList.SetCapacitiesAfterIndex(14, 1)
	// fmt.Println(bucketList)
}

func TestTryToResizeAtIndex(t *testing.T) {
	bucketList.ResetCapacities()
	bucketList.SetCapacities(7)
	itemToAdd := "G"
	bucketList.AddItemAtIndex(0, itemToAdd)
	bucketList.AddItemAtIndex(0, itemToAdd)
	bucketList.AddItemAtIndex(0, itemToAdd)

	// fmt.Println(bucketList)
	// err := bucketList.AddItemAtIndex(0, itemToAdd)
	// fmt.Println(err)

	bucketList.AddItemAtIndex(1, itemToAdd)
	bucketList.AddItemAtIndex(1, itemToAdd)
	// err2 := bucketList.AddItemAtIndex(1, itemToAdd)

	// fmt.Println(bucketList)
	// fmt.Println(err2)
}

func TestAddBucket(t *testing.T) {
	var bucketListEmpty = BucketList{make([]Bucket, 0)}
	var bucketToAdd = Bucket{"GISi", make([]string, 0), 0, 0}

	bucketListEmpty.AddBucket(bucketToAdd)
	// fmt.Println(bucketListEmpty)
}

func TestGetAvailableCapacityBeforeIndex(t *testing.T) {
	var bl = BucketList{make([]Bucket, 0)}
	var bucketToAdd = Bucket{"GISi", []string{"G", "I"}, 3, 1}
	var bucketToAdd1 = Bucket{"GISi", []string{"G", "I"}, 3, 1}
	var bucketToAdd2 = Bucket{"GISi", []string{"G", "I"}, 3, 1}

	bl.AddBucket(bucketToAdd)
	bl.AddBucket(bucketToAdd1)
	bl.AddBucket(bucketToAdd2)

	fmt.Println(bl.GetAvailableCapacityBeforeIndex(2)) //should be 2
}
