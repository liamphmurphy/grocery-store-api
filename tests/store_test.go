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
		Name:  "testing",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 0.0,
	}

	testStore.AddProduce(expectedProduce) // test the method

	// confirm that the produce item was created and it was the produce item that we expect
	if len(testStore.ProduceItems) != 1 || !testStore.ProduceItems[0].Compare(expectedProduce) {
		fmt.Println("The first produce item is:", testStore.ProduceItems[0])
		t.Errorf("The test store does not have the state that it should have")
	}
}

func TestFindProduce(t *testing.T) {
	testStore := produce_api.CreateStore()

	expectedProduce := produce_api.Produce{
		Name:  "testing",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 0.0,
	}
	testStore.ProduceItems = append(testStore.ProduceItems, expectedProduce) // add to store's internal DB

	index, produceItem := testStore.FindProduce("ABCD-1234-EFGH-5678")
	// if the index is -1 or the code is not correct, then something went wrong
	if index < 0 || produceItem.Code != "ABCD-1234-EFGH-5678" {
		t.Errorf("There was an error searching for the produce item, something is wrong")
	}
}

func TestRemoveProduce(t *testing.T) {
	testStore := produce_api.CreateStore()

	placeholder := produce_api.Produce{
		Name:  "testing",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 0.0,
	}
	// add the placeholder to the internal slice for testStore
	testStore.ProduceItems = append(testStore.ProduceItems, placeholder)

	tempList, err := testStore.RemoveProduce(placeholder.Code)

	if err != nil || len(tempList) != 0 {
		t.Errorf("It seems like the produce item still exists, something went wrong")
	}
}
