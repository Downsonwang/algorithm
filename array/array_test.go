package array

import (
	"testing"
)

func TestInsertNumIntoDataInIndex(t *testing.T) {
	cap := 10
	arr := NewArray(cap)

	for i := 0; i < cap-2; i++ {
		arr.InsertNumIntoDataInIndex(i, i+1)
	}
	arr.Print()

	arr.InsertNumIntoDataInIndex(100, 5)
	arr.Print()
}

func TestFind(t *testing.T) {
	cap := 10
	arr := NewArray(cap)
	for i := 0; i < cap; i++ {
		arr.InsertNumIntoDataInIndex(i, i+1)
	}
	arr.Print()

	arr.FindDataByIndex(4)
	arr.Print()
}

func TestInsertToTail(t *testing.T) {
	cap := 2
	arr := NewArray(cap)
	for i := 0; i < cap; i++ {
		arr.InsertNumIntoDataInIndex(i, i+1)
	}
	arr.Print()

	arr.InsertToTail()
	arr.Print()
}

func TestDeleteNumAtIndexInData(t *testing.T) {
	cap := 2
	arr := NewArray(cap)
	for i := 0; i < cap; i++ {
		arr.InsertNumIntoDataInIndex(i, i+1)
	}
	arr.Print()

	arr.DeleteNumAtIndexInData(0)
	arr.Print()
}
