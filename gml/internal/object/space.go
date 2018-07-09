package object

import (
	m "github.com/silbinarywolf/gml-go/gml/internal/math"
)

const (
	spaceBucketSize = 256
)

type SpaceObject struct {
	*Space
	spaceIndex int
}

func (space *SpaceObject) SpaceIndex() int {
	return space.spaceIndex
}

type Space struct {
	m.Vec        // Position (contains X,Y)
	Size  m.Size // Size (X,Y)
}

type SpaceBucketArray struct {
	length  int
	buckets []*SpaceBucket
}

type SpaceBucket struct {
	usedCount int
	spaces    [spaceBucketSize]Space
	used      [spaceBucketSize]bool
}

func NewSpaceBucketArray() *SpaceBucketArray {
	result := new(SpaceBucketArray)
	return result
}

func (array *SpaceBucketArray) GetNew() int {
	for b, bucket := range array.buckets {
		index := bucket.getNew()
		if index == -1 {
			continue
		}
		array.length++
		return index + (b * spaceBucketSize)
	}
	// Create new bucket, all other buckets are full!
	bucket := new(SpaceBucket)
	b := len(array.buckets)
	array.buckets = append(array.buckets, bucket)
	array.length++
	return bucket.getNew() + (b * spaceBucketSize)
}

func (array *SpaceBucketArray) Get(index int) *Space {
	bucket := array.buckets[index/spaceBucketSize]
	return &bucket.spaces[index%spaceBucketSize]
}

func (array *SpaceBucketArray) Remove(index int) {
	bucket := array.buckets[index/spaceBucketSize]
	bucketIndex := index % spaceBucketSize
	if !bucket.used[bucketIndex] {
		panic("Invalid operation. Cannot remove unused Space{} object.")
	}
	bucket.remove(bucketIndex)
	array.length--
}

func (array *SpaceBucketArray) Buckets() []*SpaceBucket {
	return array.buckets
}

// NOTE(Jake): 2018-07-08
//
// Experimented with returning this as a second value in `Get`, however
// the thing that yielded the best performance on both JS and native outputs
// was checking "IsUsed()" seperately.
//
// In fact, checking "IsUsed" seems to be an almost-free operation when added
// to the end of the if-statement checking collisions
//
func (array *SpaceBucketArray) IsUsed(index int) bool {
	bucket := array.buckets[index/spaceBucketSize]
	return bucket.used[index%spaceBucketSize]
}

func (array *SpaceBucketArray) Len() int {
	return array.length
}

func (bucket *SpaceBucket) Get(index int) *Space {
	return &bucket.spaces[index]
}

func (bucket *SpaceBucket) IsUsed(index int) bool {
	return bucket.used[index]
}

func (_ *SpaceBucket) Len() int {
	return spaceBucketSize
}

func (bucket *SpaceBucket) remove(index int) {
	bucket.used[index] = false
	bucket.usedCount--
}

func (bucket *SpaceBucket) getNew() int {
	if bucket.usedCount == spaceBucketSize {
		return -1
	}
	for i, _ := range bucket.used {
		if !bucket.used[i] {
			bucket.used[i] = true
			bucket.usedCount++
			return i
		}
	}
	return -1
}
