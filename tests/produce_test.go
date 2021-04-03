package test

import (
	"testing"

	"github.com/murnux/grocery-store-api/produce_api"
)

func TestCreateProduce(t *testing.T) {
	// create expected struct
	expectedProduce := produce_api.Produce{
		Name:  "testing_123",
		Code:  "ABCD-1234-EFGH-5678",
		Price: 12.34,
	}

	// test the CreateProduce method
	testProduce, _ := produce_api.CreateProduce("testing_123", "ABCD-1234-EFGH-5678", 12.34)

	if !expectedProduce.Compare(testProduce) {
		t.Errorf("The test produce struct does not match the expected produce struct.")
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
