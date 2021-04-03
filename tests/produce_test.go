package test

import (
	"testing"

	"github.com/murnux/grocery-store-api/produce_api"
)

// table driven tests for CreateProduce
func TestCreateProduct(t *testing.T) {
	// expected produce for the valid produce test
	expectedProduce := produce_api.Produce{
		Name:  "Test Produce",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 12.34,
	}

	tests := map[string]struct {
		name       string
		code       string
		price      float64
		expected   produce_api.Produce
		inputValid bool
	}{
		"create valid produce": {name: "Test Produce", code: "abcd-1234-EFGH-5678", price: 12.34, expected: expectedProduce, inputValid: true},
		"nil produce":          {name: "Test Invalid Produce", code: "ABC-123-343-bfe", price: 12.34, expected: produce_api.Produce{}, inputValid: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			newProduce, err := produce_api.CreateProduce(test.name, test.code, test.price)

			// conditions for error when the input is valid
			if test.inputValid {
				if !newProduce.Compare(test.expected) {
					t.Errorf("Expected: %v but got: %v", test.expected, newProduce)
				} else if err != nil {
					t.Errorf("Ran into the following error: %v", err)
				}
			} else { // conditions for error when the input is NOT valid
				if err == nil {
					t.Errorf("An error should have been returned, but the error is nil.")
				}
			}
		})
	}
}

// table driven test for testing the IsValid function
func TestIsValid(t *testing.T) {
	// create a map for the tests, with the key being the name of the test
	tests := map[string]struct {
		input      produce_api.Produce
		inputValid bool // indicates whether the input is designed to be valid or not.
	}{
		"valid produce":  {input: produce_api.Produce{Name: "Testing", Code: "ABCD-1234-EFGH-5678", Price: 1.24}, inputValid: true},
		"invalid code":   {input: produce_api.Produce{Name: "InvalidCode", Code: "ABCD-12-EF-5678", Price: 1.24}, inputValid: false},
		"negative price": {input: produce_api.Produce{Name: "InvalidPRice", Code: "ABCD-1234-EFGH-5678", Price: -1.00}, inputValid: false},
	}

	// run tests
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := produce_api.IsValid(test.input)

			// check that err is not nil
			if test.inputValid {
				if err != nil {
					t.Errorf("IsValid believes that this produce struct is invalid, something is wrong.")
				}
			} else { // when input is designed not to be valid, check if err is nil
				if err == nil {
					t.Errorf("IsValid believes this is a valid produce struct, something is wrong.")
				}
			}
		})
	}
}
