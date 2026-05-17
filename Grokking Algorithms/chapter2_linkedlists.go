package main

import (
	"errors"
	"fmt"
)

// LinkedListItem represents a node in the linked list
type LinkedListItem struct {
	value       any
	nextAddress *LinkedListItem
}

// LinkToItem connects the current node to a target node
func (l *LinkedListItem) LinkToItem(pointer *LinkedListItem) (bool, error) {
	if l == nil {
		return false, errors.New("receiver is nil")
	}
	if pointer == nil {
		return false, errors.New("the provided item is a nil pointer")
	}

	l.nextAddress = pointer
	return true, nil
}

// NextItem returns the pointer to the next node
func (l *LinkedListItem) NextItem() (*LinkedListItem, error) {
	if l == nil {
		return nil, errors.New("receiver is nil")
	}

	if l.nextAddress == nil {
		return nil, errors.New("end of list: next item is nil")
	}

	return l.nextAddress, nil
}

func main() {
	fmt.Println("Testing linkedlists")

	item1 := LinkedListItem{
		value:       "Item1",
		nextAddress: nil,
	}

	fmt.Println("Item 1:", item1)
	fmt.Printf("Raw pointer Item 1: %p\n\n", &item1)

	item2 := LinkedListItem{
		value:       "Item2",
		nextAddress: nil,
	}
	fmt.Printf("Raw pointer Item 2: %p\n\n", &item2)

	// Link the items
	item1.LinkToItem(&item2)
	fmt.Println("Item 1 after linking:", item1)

	// Retrieve and print the next item
	nextItem, err := item1.NextItem()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nNext item raw pointer: %p\n", nextItem)
		fmt.Println("Next item content:", nextItem)
	}
}