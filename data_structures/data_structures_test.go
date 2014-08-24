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

// func TestAddItemAtIndex(t *testing.T) {
// 	bucketList.ResetCapacities()
// 	bucketList.SetCapacities(5)
// 	itemToAdd := "G"
// 	bucketList.AddItemAtIndex(0, itemToAdd)
// 	bucketList.AddItemAtIndex(0, itemToAdd)
// 	bucketList.AddItemAtIndex(1, itemToAdd)
// 	fmt.Println(bucketList)
// }

func TestTryToResizeAtIndex(t *testing.T) {
	bucketList.ResetCapacities()
	bucketList.SetCapacities(7)
	itemToAdd := "G"
	bucketList.AddItemAtIndex(0, itemToAdd)
	bucketList.AddItemAtIndex(0, itemToAdd)
	bucketList.AddItemAtIndex(0, itemToAdd)

	fmt.Println(bucketList)
	// err := bucketList.TryToResizeAtIndex(0)
	// fmt.Println(err)
	// fmt.Println(bucketList)
	err := bucketList.AddItemAtIndex(0, itemToAdd)
	fmt.Println(err)

}
