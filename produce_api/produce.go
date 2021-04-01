package produce_api

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

// produce represents the structure of a single produce item
type Produce struct {
	Name  string  `json:"Name"`
	Code  string  `json:"Code"`
	Price float64 `json:"Price"`
}

// CreateProduce creates a single produce item with passed in values
func CreateProduce(name string, code string, price float64) (Produce, error) {
	var new_produce Produce

	// assign struct values
	new_produce.Name = name
	new_produce.Code = strings.ToUpper(code)          // always ensure the produce code is upper case
	new_produce.Price = math.Round(price/0.01) * 0.01 // round to two decimal places

	// make sure the produce struct is valid
	err := IsValid(new_produce)
	if err != nil {
		return new_produce, err
	}

	return new_produce, nil
}

// Compare checks whether the current Produce struct and another Produce struct is equal based on attributes
func (currentProduce Produce) Compare(otherProduce Produce) bool {
	// check memory address first
	if &currentProduce == &otherProduce {
		return true
	}

	// check attributes, could do this all in one line but that would look gross
	if currentProduce.Name != otherProduce.Name {
		return false
	} else if currentProduce.Code != otherProduce.Code {
		return false
	} else if currentProduce.Price != otherProduce.Price {
		return false
	}

	// if we got this far, then the two structs should be the same
	return true
}

// IsValid confirms that the passed in produce is valid, for example that the produce code is of a valid format.
func IsValid(produce Produce) error {
	// check regexp matches desired format for the produce code
	matched, _ := regexp.MatchString("[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}", produce.Code)

	// cycle through validity checks
	if !matched {
		return fmt.Errorf("the produce code: %v is not valid", produce.Code)
	} else if produce.Price < 0.0 {
		return fmt.Errorf("the product price %v is negative, please use a positive floating point number", produce.Price)
	}
	return nil
}
