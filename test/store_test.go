package test

import (
	"fmt"
	"testing"

	"github.com/murnux/grocery-store-api/produce_api"
)

func TestCreateEmptyStore(t *testing.T) {
	// make a basic store which should be empty
	testStore := produce_api.CreateStore()

	if len(testStore.ProduceItems) != 0 {
		t.Errorf("The test store is deemed not correct as it is not empty.")
	}
}

func TestAddProduce(t *testing.T) {
	testStore := produce_api.CreateStore()

	expectedProduce := produce_api.Produce{
		Name:        "testing",
		ProduceCode: "ABCD-1234-EFGH-5678",
		Price:       0.0,
	}

	testStore.AddProduce(expectedProduce) // test the method

	if len(testStore.ProduceItems) != 1 || !testStore.ProduceItems[0].Compare(expectedProduce) {
		fmt.Println("The first produce item is:", testStore.ProduceItems[0])
		t.Errorf("The test store does not have the state that it should have")
	}
}
