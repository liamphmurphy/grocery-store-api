package main

import (
	"github.com/murnux/grocery-store-api/produce_api"
)

func main() {
	store := produce_api.CreateStore() // create store struct model for use in the API
	store.PopulateDefaultProduce()     // populate default produce items as specified in the specifications

	produce_api.APIMain(store) // start the API (TODO: determine if goroutine is needed)
}
