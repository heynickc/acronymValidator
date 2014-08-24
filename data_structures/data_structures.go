package data_structures

import (
	"errors"
)

type Bucket struct {
	Name              string
	Items             []string
	Capacity          int
	AvailableCapacity int
}

func (b *Bucket) AddItem(item string) error {
	if b.AvailableCapacity > 0 {
		b.Items = InsertStringSlice(b.Items, []string{item}, 0)
		b.AvailableCapacity--
		return nil
	} else {
		return errors.New("Can't add anymore")
	}
}

type BucketList struct {
	Buckets []Bucket
}

func (bl *BucketList) AddItemAtIndex(index int, item string) error {
	err := bl.Buckets[index].AddItem(item)
	if err != nil {
		err := bl.TryToResizeAtIndex(index)
		if err != nil {
			return errors.New("Can't resize")
		}
		bl.AddItemAtIndex(index, item)
	}
	return nil
}

func (bl *BucketList) AddBucket(bucket Bucket) {
	bl.Buckets = append(bl.Buckets, bucket)
}

func (bl *BucketList) TryToResizeAtIndex(index int) error {
	if len(bl.Buckets[index+1:]) >= bl.GetTotalCapacityAfterIndex(index)-1 {
		return errors.New("Not enough Capacity left")
	} else {
		bl.Buckets[index].Capacity++
		bl.Buckets[index].AvailableCapacity++
		bl.SetCapacitiesAfterIndex(bl.GetTotalCapacityAfterIndex(index)-1, index)
	}
	return nil
}

func (bl *BucketList) ResetCapacities() {
	for i := 0; i < len(bl.Buckets); i++ {
		bl.Buckets[i].Capacity = 0
		bl.Buckets[i].AvailableCapacity = 0
	}
}

func (bl *BucketList) ResetCapacitiesAfterIndex(index int) {
	for i := index + 1; i < len(bl.Buckets); i++ {
		bl.Buckets[i].Capacity = 0
		bl.Buckets[i].AvailableCapacity = 0
	}
}

func (bl *BucketList) SetCapacities(count int) {
	for {
		if count == 0 {
			break
		}
		for i := 0; count > 0 && i < len(bl.Buckets); i, count = i+1, count-1 {
			bl.Buckets[i].Capacity++
			bl.Buckets[i].AvailableCapacity++
		}
	}
}

func (bl *BucketList) SetCapacitiesAfterIndex(count int, index int) {
	bl.ResetCapacitiesAfterIndex(index)

	for {
		if count == 0 {
			break
		}
		for i := index + 1; count > 0 && i < len(bl.Buckets); i, count = i+1, count-1 {
			bl.Buckets[i].Capacity++
			bl.Buckets[i].AvailableCapacity++
		}
	}
}

func (bl *BucketList) GetTotalCapacity() int {
	totalCapacity := 0
	for _, bucket := range bl.Buckets {
		totalCapacity = totalCapacity + bucket.Capacity
	}
	return totalCapacity
}

func (bl *BucketList) GetTotalCapacityAfterIndex(index int) int {
	totalCapacity := 0
	for _, bucket := range bl.Buckets[index+1:] {
		totalCapacity = totalCapacity + bucket.Capacity
	}
	return totalCapacity
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}
