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

func APIMain() {
	router := gin.Default()
	router.Use(cors.Default())

	store := CreateStore() // create store struct for use in the API
	store.PopulateDefaultProduce()

	// API GET endpoints
	router.GET("/produce/getall", store.getAllHandler)

	router.Run()
}
