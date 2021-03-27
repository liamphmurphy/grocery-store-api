package produce_api

// produce represents the structure of a single produce item
type Produce struct {
	Name        string
	ProduceCode string
	Price       float64
}

// CreateProduce creates a single produce item with passed in values
func CreateProduce(name string, code string, price float64) Produce {
	var new_produce Produce

	// assign struct values
	new_produce.Name = name
	new_produce.ProduceCode = code
	new_produce.Price = price

	return new_produce
}

// compare checks whether the current Produce struct and another Produce struct is equal based on attributes
func (currentProduce Produce) compare(otherProduce Produce) bool {
	// check memory address first
	if &currentProduce == &otherProduce {
		return true
	}

	// check attributes, could do this all in one line but that would look gross
	if currentProduce.Name != otherProduce.Name {
		return false
	} else if currentProduce.ProduceCode != otherProduce.ProduceCode {
		return false
	} else if currentProduce.Price != otherProduce.Price {
		return false
	}

	// if we got this far, then the two structs should be the same
	return true
}
