// This file will hold a slice of produce and add functionality for a store / collection of produce
// that wouldn't make sense in produce.go itself. I am assuming this would help with adding multiple stores / chains
// if this was a real world product. For our purposes, you can think of this file as the 'database'

package produce_api

import "errors"

type Store struct {
	ProduceItems []Produce
}

func CreateStore() *Store {
	// create empty slice of Produce items
	var produceItems []Produce

	var store Store
	store.ProduceItems = produceItems
	return &store
}

func (store *Store) AddProduce(newItem Produce) error {
	initialSize := len(store.ProduceItems) // used for comparing length after appending later

	store.ProduceItems = append(store.ProduceItems, newItem)
	if len(store.ProduceItems) == initialSize+1 {
		return nil // indicate a successful append
	} else {
		return errors.New("the state of the list of produce items is still the same after appending, something went wrong")
	}
}
