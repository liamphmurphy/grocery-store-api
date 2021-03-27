package produce_api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// getaAllHandler returns the JSON of all produce items
func (store *Store) getAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, store.ProduceItems)
}

// getProduceHandler returns the JSON of one or more produce items based on URL paramaters
// expects a 'code' parameter containing the produce code
func (store *Store) getProduceHandler(c *gin.Context) {
	params := c.Request.URL.Query()

	// create a new slice to hold all of the returned produce items
	var foundProduce []Produce
	for _, code := range params["code"] {
		index, produceItem := store.FindProduce(code)
		if index >= 0 { // if no error returned, then assume the product is valid
			foundProduce = append(foundProduce, produceItem)
		}
	}

	if len(foundProduce) > 0 { // return produce data if any was found
		c.JSON(http.StatusOK, foundProduce)
	} else {
		// return that the request was processed, but no data was found
		c.JSON(http.StatusNoContent, gin.H{"Warning": "No produce items matched the provided produce code(s)."})
	}
}

// addProduceHandler handles the POST request from a client for adding a produce item to the internal DB
func (store *Store) addProduceHandler(c *gin.Context) {

}

// deleteProduceHandler handles the DELETE request when a client requests to delete a produce item
func (store *Store) deleteProduceHandler(c *gin.Context) {

}

func APIMain() {
	router := gin.Default()
	router.Use(cors.Default())

	store := CreateStore() // create store struct for use in the API
	store.PopulateDefaultProduce()

	// API GET endpoints
	router.GET("/produce/getall", store.getAllHandler)
	router.GET("/produce/getitem", store.getProduceHandler)

	// API POST endpoints
	router.GET("/produce/addproduce", store.addProduceHandler)

	router.Run()
}
