package data_structures

import (
	"fmt"
)

type Bucket struct {
	name     string
	items    []string
	capacity int
}

type BucketList struct {
	buckets []Bucket
}

func (bl *BucketList) SetCapacities(count int) {
	for {
		if count == 0 {
			break
		}
		for i := 0; count > 0 && i < len(bl.buckets); i, count = i+1, count-1 {
			bl.buckets[i].capacity++
		}
	}
}
