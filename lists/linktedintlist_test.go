package linkedintlist

import "testing"

func TestCreate(t *testing.T) {
	var ll = new()
	if !ll.isEmpty() {
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	var ll = new()
	ll.append(1)
	if ll.isEmpty() {
		t.Errorf("List is empty after append.")
	}
}

func TestLen(t * testing.T) {
	var ll = new()
	if ll.len() != 0 {
		t.Errorf("Empty list has non-zero length.")
	}
	ll.append(1)
	if ll.len() != 1 {
		t.Errorf("Single element list has non-one length.")
	}
}

func TestGet(t *testing.T) {
	var ll = new()
	ll.append(1)
	ll.append(2)

	val, _ := ll.get(0)	
	if val != 1 {
		t.Errorf("Index 0 is not 1")
	}
	val, _ = ll.get(1)
	if val != 2 {
		t.Errorf("Index 1 is not 2")
	}
	val, err := ll.get(2)
	if !err {
		t.Errorf("Should have gotten out of bounds error, instead got value %v", val)
	}
}

func TestPrepend(t *testing.T) {
	ll := new()
	ll.prepend(1)
	ll.prepend(2)
	ll.prepend(3)
	val, _ := ll.get(0) 
	if (val != 3) {
		t.Errorf("First node should contain value 3")
	}
	val, _ = ll.get(1)
	if (val != 2) {
		t.Errorf("Second node should contain value 2")
	}
	val, _ = ll.get(2) 
	if (val != 1) {
		t.Errorf("Third node should contain value 1")
	}
}

func TestPrint(t *testing.T) {
	var ll = new()
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.print()
}

func TestInsert(t *testing.T) {
	var ll = new()
	ll.append(1)
	ll.append(3)
	ll.insert(2, 1)
	ll.insert(4, 3)
	val, _ := ll.get(0)	
	if val != 1 {
		t.Errorf("Index 0 is not 1")
	}
	val, _ = ll.get(1)
	if val != 2 {
		t.Errorf("Index 1 is not 2: %v", val)
	}
	val, _ = ll.get(2)	
	if val != 3 {
		t.Errorf("Index 2 is not 3: %v", val)
	}
	val, _ = ll.get(3)
	if val != 4 {
		t.Errorf("Index 3 is not 4: %v", val)
	}
	ll.print()
}

func TestDelete(t *testing.T) {
	var ll = new()
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.delete(4)
	ll.print()
	ll.delete(0)
	ll.print()
	ll.delete(1)
	ll.print()
	ll.delete(0)
	ll.print()
	ll.delete(0)
	ll.print()
}