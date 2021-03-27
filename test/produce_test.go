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

func TestIsValid(t *testing.T) {
	// create test produce struct that has valid values
	produce := produce_api.CreateProduce("Test", "ABC1-DEF2-GHI3-JKL4", 10.00)

	err := produce_api.IsValid(produce)
	if err != nil {
		t.Errorf("IsValid believes that this produce struct is invalid, something is wrong.")
	}
}
