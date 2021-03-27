package test

import (
	"testing"

	"github.com/murnux/grocery-store-api/produce_api"
)

func TestCreateProduce(t *testing.T) {
	// create expected struct
	expectedProduce := produce_api.Produce{
		Name:        "testing_123",
		ProduceCode: "ABCD-1234-EFGH-5678",
		Price:       12.34,
	}

	// test the CreateProduce method
	testProduce := produce_api.CreateProduce("testing_123", "ABCD-1234-EFGH-5678", 12.34)

	if !expectedProduce.Compare(testProduce) {
		t.Errorf("The test produce struct does not match the expected produce struct.")
	}
}
