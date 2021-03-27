package main

import (
	"fmt"

	"github.com/murnux/grocery-store-api/produce_api"
)

// this is at a "test driver" stage, still determining the best way to interface with the api
func main() {
	fmt.Println("test")

	produce_api.APIMain() // start the API (TODO: determine if goroutine is needed)
}
