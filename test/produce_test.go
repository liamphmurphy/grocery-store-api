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
	testProduce, _ := produce_api.CreateProduce("testing_123", "ABCD-1234-EFGH-5678", 12.34)

	if !expectedProduce.Compare(testProduce) {
		t.Errorf("The test produce struct does not match the expected produce struct.")
	}
}

func TestIsValid(t *testing.T) {
	// create test produce struct that has valid values
	produce := produce_api.Produce{
		Name:        "testing_123",
		ProduceCode: "ABCD-1234-EFGH-5678",
		Price:       12.34,
	}

	err := produce_api.IsValid(produce)
	if err != nil {
		t.Errorf("IsValid believes that this produce struct is invalid, something is wrong.")
	}
}

// This is the opposite of TestIsValid; purposefully passes in an invalid produce cod
func TestIsNotValidCode(t *testing.T) {
	// create produce item with an invalid code
	produce := produce_api.Produce{
		Name:        "testing_123",
		ProduceCode: "ThisIsNotAValidCode",
		Price:       12.34,
	}

	err := produce_api.IsValid(produce)
	if err == nil { // err should not be nil, so check if it is
		t.Errorf("IsValid believes this is a valid produce struct, something is wrong.")
	}
}

// makes sure IsValid catches a negative value price
func TestIsNotValidPrice(t *testing.T) {
	// create produce item with an invalid price
	produce := produce_api.Produce{
		Name:        "testing_123",
		ProduceCode: "ABCD-1234-EFGH-5678",
		Price:       -1.00,
	}

	err := produce_api.IsValid(produce)
	if err == nil {
		t.Errorf("IsValid believes this is a valid produce struct, something is wrong.")
	}
}
