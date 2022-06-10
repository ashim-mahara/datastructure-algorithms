package main

import (
	"fmt"
	"testing"
)

var ll = linkedList{}

func init() {
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.append(6)
	ll.append(7)
	ll.append(8)

	fmt.Printf("The following linkedList was created for the tests: \n\n")
	ll.printLinkedList()
	fmt.Printf("\n")
}

func TestGetNode(t *testing.T) {
	want := 4
	node, _ := ll.getNode(4)

	got := node.value

	if got != want {
		t.Errorf("Error in getNode. Wanted %v, git %v", want, got)
	}
}

func TestUpdate(t *testing.T) {
	update_position, update_value := 4, 12
	ll.update(update_position, update_value)

	node, _ := ll.getNode(4)

	got := node.value

	if got != 12 {
		t.Errorf("Error in update function. Wanted %v, got %v.", update_value, got)
	}

}

func TestRemove(t *testing.T) {
	remove_position := 6
	ll.remove(remove_position)

	node, _ := ll.getNode(6)

	got := node.value

	if got != 7 {
		t.Errorf("Error is remove function. Wanted %v, got %v.", 7, got)
	}
}

func TestInsert(t *testing.T) {
	insert_position := 4
	insert_value := 61

	ll.insert(insert_position, insert_value)
	node, _ := ll.getNode(insert_position)

	got := node.value
	if got != insert_value {
		t.Errorf("Error in insert function. Wanted %v, got %v", insert_value, got)
		ll.printLinkedList()
	}
}

func TestPrepend(t *testing.T) {
	prepend_value := 0

	prev_head := ll.head.value

	ll.prepend(prepend_value)

	got := ll.head.value
	got_next := ll.head.next.value

	if got != 0 {
		t.Errorf("Error in prepend function. Wanted head value as %v, got %v", prepend_value, got)
		ll.printLinkedList()
	} else if got_next != prev_head {
		t.Errorf("Error in prepend function. Wanted head.next.value as %v, got %v", prev_head, got_next)
		ll.printLinkedList()
	}
}

func TestPop(t *testing.T) {
	ll.pop()

	node, _ := ll.getNode(-1)
	got := node.value

	if got != 7 {
		t.Errorf("Error in the pop function. Wanted %v, got %v", 7, got)
	}
}

func TestHead(t *testing.T) {
	if ll.head.value != 0 {
		t.Errorf("Error while assigning values to ll.head. Wanted 1, got %v", ll.head)
	}
}

func TestTail(t *testing.T) {
	tailValue := ll.tail.value
	if tailValue != 7 {
		t.Errorf("Error while assigning values to ll.tail. Wanted 7, got %v", tailValue)
	}
}

func TestLength(t *testing.T) {
	if ll.elem != 8 {
		t.Errorf("Error while getting the length of linkedList. Wanted 8, got %v", ll.elem)
		ll.printLinkedList()
	}
}