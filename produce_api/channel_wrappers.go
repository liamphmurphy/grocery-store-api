// this file contains functions using channels that are wrappers around functions written in store.go and produce.go
// the idea is to separate the logic of these other functions (e.g FindProduce, RemoveProduce, etc) and the logic to support channels.
// while this takes some more work, it allows users to use these functions with or without goroutines

package produce_api

func (store *Store) FindProduceChannel(code string, produceChan chan Produce, indexChan chan int) {
	index, foundProduce := store.FindProduce(code)
	if index != -1 {
		produceChan <- foundProduce // assign found produce item to the channel
	}
	indexChan <- index // assign index to the index channel
}
