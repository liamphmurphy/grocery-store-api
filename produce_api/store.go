// This file will hold a slice of produce and add functionality for a store / collection of produce
// that wouldn't make sense in produce.go itself. I am assuming this would help with adding multiple stores / chains
// if this was a real world product. For our purposes, you can think of this file as the 'database'

package produce_api

import "errors"

// for now I'll be going with a slice over a map
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
	// iterate over internal slice
	for index, produceItem := range store.ProduceItems {
		if produceItem.ProduceCode == code {
			return index, produceItem
		}
	}

	var item Produce // create empty produce item for sake of the return statement
	return -1, item
}

// RemoveProduce takes in a code and removes the associated produce item from the internal DB
func (store *Store) RemoveProduce(code string) []Produce {
	// perform a standard swap and resize of the slize
	// TODO: this method is going to get slow since we have to find the produce item first, this may prompt me to use a map later on
	index, _ := store.FindProduce(code)
	// perform the swap
	var temp []Produce
	copy(store.ProduceItems, temp)
	temp[len(temp)-1], temp[index] = temp[index], temp[len(temp)-1]
	return temp
}
