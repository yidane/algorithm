package problems

import (
	"testing"
)

func TestMyCircularDequeConstructor(t *testing.T) {

	obj := MyCircularDequeConstructor(3)
	param1 := obj.InsertFront(1)
	if !param1 {
		t.Fatal()
	}
	param2 := obj.InsertLast(2)
	if !param2 {
		t.Fatal()
	}
	param3 := obj.DeleteFront()
	if !param3 {
		t.Fatal()
	}
	param4 := obj.DeleteLast()
	if !param4 {
		t.Fatal()
	}
	param5 := obj.GetFront()
	if param5 != -1 {
		t.Fatal()
	}
	param6 := obj.GetRear()
	if param6 != -1 {
		t.Fatal()
	}
	param7 := obj.IsEmpty()
	if !param7 {
		t.Fatal()
	}
	param8 := obj.IsFull()
	if param8 {
		t.Fatal()
	}
}
