package main

import (
	"fmt"

	"github.com/murnux/grocery-store-api/produce_api"
)

// this is at a "test driver" stage, still determining the best way to interface with the api
func main() {
	fmt.Println("test")

	product := produce_api.CreateProduce("test", "test", 0.00)

	fmt.Println(product)
}
