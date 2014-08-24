package data_structures

import (
	"fmt"
	// "errors"
)

type Bucket struct {
	name              string
	items             []string
	capacity          int
	availableCapacity int
}

func (b *Bucket) AddItem(item string) bool {
	if b.availableCapacity > 0 {
		b.items = InsertStringSlice(b.items, []string{item}, 0)
		b.availableCapacity--
		return true
	} else {
		return false
	}
}

type BucketList struct {
	buckets []Bucket
}

func (bl *BucketList) AddItemAtIndex(index int, item string) {
	ok := bl.buckets[index].AddItem(item)
	if !ok {
		fmt.Println("Not ok")
		fmt.Println(bl.buckets[index].availableCapacity)
	}
}

func (bl *BucketList) TryToResizeAtIndex(index int) {

}

func (bl *BucketList) ResetCapacities() {
	for i := 0; i < len(bl.buckets); i++ {
		bl.buckets[i].capacity = 0
		bl.buckets[i].availableCapacity = 0
	}
}

func (bl *BucketList) ResetCapacitiesAfterIndex(index int) {
	for i := index; i < len(bl.buckets); i++ {
		bl.buckets[i].capacity = 0
		bl.buckets[i].availableCapacity = 0
	}
}

func (bl *BucketList) SetCapacities(count int) {
	for {
		if count == 0 {
			break
		}
		for i := 0; count > 0 && i < len(bl.buckets); i, count = i+1, count-1 {
			bl.buckets[i].capacity++
			bl.buckets[i].availableCapacity++
		}
	}
}

func (bl *BucketList) SetCapacitiesAfterIndex(count int, index int) {
	bl.ResetCapacitiesAfterIndex(index)

	for {
		if count == 0 {
			break
		}
		for i := index; count > 0 && i < len(bl.buckets); i, count = i+1, count-1 {
			bl.buckets[i].capacity++
			bl.buckets[i].availableCapacity++
		}
	}
}

func (bl *BucketList) GetTotalCapacity() int {
	totalCapacity := 0
	for _, bucket := range bl.buckets {
		totalCapacity = totalCapacity + bucket.capacity
	}
	return totalCapacity
}

func (bl *BucketList) GetTotalCapacityAfterIndex(index int) int {
	totalCapacity := 0
	for _, bucket := range bl.buckets[index+1:] {
		totalCapacity = totalCapacity + bucket.capacity
	}
	return totalCapacity
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}
