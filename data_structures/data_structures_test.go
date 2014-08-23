package data_structures

import (
	"fmt"
	"testing"
)

func TestSetCapacity(t *testing.T) {
	bucket := Bucket{"TEST", []string{"item", "item2", "item3"}, 0}
	bucket1 := Bucket{"TEST", []string{"item", "item2", "item3"}, 0}
	bucket2 := Bucket{"TEST", []string{"item", "item2", "item3"}, 0}
	bucket3 := Bucket{"TEST", []string{"item", "item2", "item3"}, 0}

	bucketList := BucketList{[]Bucket{bucket, bucket1, bucket2, bucket3}}

	bucketList.SetCapacities(15)
	fmt.Println(bucketList)
}
