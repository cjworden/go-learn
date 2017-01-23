package linkedintlist

import (
	"fmt"
	"errors"
)

type node struct {
	prev *node
	next *node
	value int
}

type LinkedIntList struct {
	head *node
	tail *node
	size int
}

func new() LinkedIntList {
	return LinkedIntList{nil, nil, 0}
}

func (ll *LinkedIntList) append (i int) {
	var n = node{ll.tail, nil, i}
	if (ll.isEmpty()) {
		ll.head = &n
		ll.tail = &n
	}else {
		ll.tail.next = &n
		ll.tail = &n
	}
	ll.size++
}

func (ll *LinkedIntList) prepend (i int) {
	var n = node{nil, ll.head, i}
	if (ll.isEmpty()) {
		ll.head = &n
		ll.tail = &n
	}
	var old = ll.head
	ll.head = &n
	old.prev = &n
	ll.size++
}

func (ll *LinkedIntList) insert (i int, index int) error {
	// OutOfBounds error
	if (index < 0 || index > ll.len()) {
		return errors.New("Index out of bounds for insert operation.")
	}
	if (index == ll.len()) {
		ll.append(i)
		return nil
	}else if (index == 0) {
		ll.prepend(i)
		return nil
	}
	// Loop to insertion point
	var curr = ll.head
	var x = 0
	for ;x < index; x++ {
		curr = curr.next
	}
	
	n := node{curr.prev, curr, i}
	curr.prev.next = &n
	curr.prev = &n
	ll.size++
	return nil
}

func (ll *LinkedIntList) delete (index int) error{
	if (index < 0 || index >= ll.len()) {
		return errors.New("Index out of bounds for delete operation.")
	}
	
	if (index == 0) {
		if (ll.len() == 1) {
			// Removing last element, head and tail = nil
			ll.head = nil
			ll.tail = nil
		} else {
			// Deleting head element requires fixing head
			ll.head.next.prev = nil
			ll.head = ll.head.next
		}
	}else if (index == ll.len() - 1) {
		// Deleting tail element requries fixing tail
		ll.tail.prev.next = nil
		ll.tail = ll.tail.prev
	} else {
		var curr = ll.head
		for x:= 0; x < index; x++ {
			curr = curr.next
		}
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}
	ll.size--
	return nil
}

func (ll *LinkedIntList) get (index int) (int, bool) {
	// Get of empty list is an error
	if ll.isEmpty() {
		return -1, true
	}
	// OutOfBounds error
	if (index < 0 || index >= ll.len()) {
		return -1, true
	}

	var curr = ll.head
	for x := 0; x < index; x++ {
		curr = curr.next
	}
	return curr.value, false
}

func (ll *LinkedIntList) isEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedIntList) len() int {
	return ll.size
}

func (ll *LinkedIntList) print() {
	fmt.Print("[")
	if ll.len() > 0 {
		printrec(*ll.head)
	}
	fmt.Println("]")
}

func printrec(n node) {
	fmt.Printf("%v", n.value)
	if n.next != nil {
		// Must double quote the , to make it a string instead of a rune...
		fmt.Print(",")
		printrec(*n.next)
	}
}