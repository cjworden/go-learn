package arrays

import "testing"

func TestCreate(t *testing.T) {
	array := Create()
	for _, x := range array {
		if x != 0 {
			t.Fail()
		}
	}
}

func TestInitialize(t *testing.T) {
	array := Initialize()
	ndx := 0
	for _, x := range array {
		if x != ndx {
			t.Errorf("Array value %d doesn't match index value %d", x, ndx)
		}
		ndx++
	}
}

func TestSlice(t *testing.T) {
	array := Initialize()
	slice := array[0:2]
	if len(slice) != 2 {
		t.Errorf("Slice has size larger than 2.")
	}
}