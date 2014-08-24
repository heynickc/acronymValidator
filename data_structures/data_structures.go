package data_structures

import (
	"errors"
	"fmt"
)

type Bucket struct {
	name              string
	items             []string
	capacity          int
	availableCapacity int
}

func (b *Bucket) AddItem(item string) error {
	if b.availableCapacity > 0 {
		b.items = InsertStringSlice(b.items, []string{item}, 0)
		b.availableCapacity--
		return nil
	} else {
		return errors.New("Can't add anymore")
	}
}

type BucketList struct {
	buckets []Bucket
}

func (bl *BucketList) AddItemAtIndex(index int, item string) error {
	err := bl.buckets[index].AddItem(item)
	if err != nil {
		err := bl.TryToResizeAtIndex(index)
		if err != nil {
			return errors.New("Can't resize")
		}
		bl.AddItemAtIndex(index, item)
	}
	return nil
}

func (bl *BucketList) TryToResizeAtIndex(index int) error {
	// fmt.Println(bl.buckets[index])
	fmt.Println(bl.GetTotalCapacityAfterIndex(index))
	fmt.Println(len(bl.buckets[index+1:]))
	if len(bl.buckets[index+1:]) >= bl.GetTotalCapacityAfterIndex(index)-1 {
		return errors.New("Not enough capacity left")
	} else {
		bl.buckets[index].capacity++
		bl.buckets[index].availableCapacity++
		bl.SetCapacitiesAfterIndex(bl.GetTotalCapacityAfterIndex(index)-1, index)
	}
	return nil
}

func (bl *BucketList) ResetCapacities() {
	for i := 0; i < len(bl.buckets); i++ {
		bl.buckets[i].capacity = 0
		bl.buckets[i].availableCapacity = 0
	}
}

func (bl *BucketList) ResetCapacitiesAfterIndex(index int) {
	for i := index + 1; i < len(bl.buckets); i++ {
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
		for i := index + 1; count > 0 && i < len(bl.buckets); i, count = i+1, count-1 {
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
