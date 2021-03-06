// This file will hold a slice of produce and add functionality for a store / collection of produce
// that wouldn't make sense in produce.go itself. I am assuming this would help with adding multiple stores / chains
// if this was a real world product. For our purposes, you can think of this file as the 'database'

package produce_api

import (
	"errors"
	"strings"
)

// for now I'll be going with a slice over a map
type Store struct {
	ProduceItems []Produce
}

// CreateStore initializes a Store struct
func CreateStore() *Store {
	// create empty slice of Produce items
	var produceItems []Produce

	var store Store
	store.ProduceItems = produceItems
	return &store
}

// AddProduce adds a new produce item to the internal slice, implements the 'add' operation
func (store *Store) AddProduce(newItem Produce) error {
	err := IsValid(newItem)
	if err != nil {
		return err // if something is wrong with the produce item, pass on the error
	}

	index, _ := store.FindProduce(newItem.Code)
	if index >= 0 { // check if there is an item with this produce code already
		return errors.New("an item with this produce code already exists")
	}

	initialSize := len(store.ProduceItems) // used for comparing length after appending later

	store.ProduceItems = append(store.ProduceItems, newItem)
	// check if the slice has increased by one element
	if len(store.ProduceItems) == initialSize+1 {
		return nil // indicate a successful append
	} else {
		return errors.New("the state of the list of produce items is still the same after appending, something went wrong")
	}
}

// PopulateDefaultProduce puts in some default files per the documentation
func (store *Store) PopulateDefaultProduce() {
	// create default produce items
	lettuce, _ := CreateProduce("Lettuce", "A12T-4GH7-QPL9-3N4M", 3.46)
	peach, _ := CreateProduce("Peach", "E5T6-9UI3-TH15-QR88", 2.99)
	greenPepper, _ := CreateProduce("Green Pepper", "YRT6-72AS-K736-L4AR", 0.79)
	galaApple, _ := CreateProduce("Gala Apple", "TQ4C-VV6T-75ZX-1RMR", 3.59)

	// populate the store
	store.AddProduce(lettuce)
	store.AddProduce(peach)
	store.AddProduce(greenPepper)
	store.AddProduce(galaApple)
}

// FindProduce searches the internal produce db for a produce item based on the produce code
func (store *Store) FindProduce(code string) (int, Produce) {
	code = strings.ToUpper(code)
	// iterate over internal slice
	for index, produceItem := range store.ProduceItems {
		if produceItem.Code == code {
			return index, produceItem
		}
	}

	var item Produce // create empty produce item for sake of the return statement, will be nil
	return -1, item
}

// RemoveProduce takes in a code and removes the associated produce item from the internal DB
func (store *Store) RemoveProduce(code string) ([]Produce, error) {
	// attempt to find the produce and its associated index
	index, _ := store.FindProduce(code)
	temp := make([]Produce, len(store.ProduceItems)) // need to allocate enough space for copy to wrok

	if index == -1 {
		return temp, errors.New("could not find the produce item")
	}
	// avoid mutating original DB until changes are successfully made
	copy(temp, store.ProduceItems)

	// perform a standard swap and resize of the slice
	temp[len(temp)-1], temp[index] = temp[index], temp[len(temp)-1]
	return temp[:len(temp)-1], nil
}
