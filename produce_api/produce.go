package produce_api

// produce represents the structure of a single produce item
type produce struct {
	Name        string
	ProduceCode string
	Price       float64
}

// CreateProduce creates a single produce item with passed in values
func CreateProduce(name string, code string, price float64) produce {
	var new_produce produce

	// assign struct values
	new_produce.Name = name
	new_produce.ProduceCode = code
	new_produce.Price = price

	return new_produce
}
