package produce_api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Needed to hold a list of produce items when accepting JSON in addProduceHandler
type ProduceList struct {
	List []Produce `json:"Produce"`
}

// ResultSet defines the structure of the JSON payload returned to users
type ResultSet struct {
	Status  string      `json:"status"`  // used to give some kind of feedback to the user
	Results interface{} `json:"results"` // let the results payload be adaptable
}

// getAllHandler returns the JSON of all produce items
func (store *Store) getAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, ResultSet{Status: "", Results: store.ProduceItems})
}

// getProduceHandler returns the JSON of one or more produce items based on URL paramaters
// expects a 'code' parameter containing the produce code
func (store *Store) getProduceHandler(c *gin.Context) {
	params := c.Request.URL.Query()

	// create a new slice to hold all of the returned produce items
	var foundProduce []Produce

	// make a produce struct channel that is big enough to contain the number of codes requested
	produceChan := make(chan Produce, len(params["Produce Code"]))
	indexChan := make(chan int)
	for _, code := range params["Produce Code"] {
		// create channels necessary for the goroutine
		go store.FindProduceChannel(code, produceChan, indexChan)
		if <-indexChan != -1 { // if no error returned, then assume the product is valid
			foundProduce = append(foundProduce, <-produceChan)
		}
	}

	// be kind, close the channels
	close(produceChan)
	close(indexChan)

	if len(foundProduce) > 0 { // return produce data if any was found
		c.JSON(http.StatusOK, ResultSet{Status: "", Results: foundProduce})
	} else {
		// return that the request was processed, but no data was found
		c.JSON(http.StatusOK, ResultSet{Status: "No data was found from the query.", Results: foundProduce})
	}
}

// addProduceHandler handles the POST request from a client for adding a produce item to the internal DB
func (store *Store) addProduceHandler(c *gin.Context) {
	var list ProduceList
	c.BindJSON(&list) // bind JSON body to Produce struct

	errChan := make(chan error)          // make error channel
	preLength := len(store.ProduceItems) // get length of slice before changes, used for a check later on
	for _, produce := range list.List {
		produce, _ = CreateProduce(produce.Name, produce.ProduceCode, produce.Price) // recreate produce item to ensure correct formats
		go store.AddProduceChannel(produce, errChan)                                 // add the new produce to the db
		if <-errChan != nil {
			c.JSON(http.StatusBadRequest, ResultSet{Status: fmt.Sprintf("Could not add the item: %v\n", produce.ProduceCode), Results: nil})
			return
		}
	}
	close(errChan)

	// check if the internal slice was adjusted at all
	if preLength == len(store.ProduceItems) {
		c.JSON(http.StatusOK, ResultSet{Status: "No items were added", Results: nil})
	} else {
		c.JSON(http.StatusAccepted, ResultSet{Status: "The item(s) have been added", Results: list})
	}
}

// deleteProduceHandler handles the DELETE request when a client requests to delete a produce item
func (store *Store) deleteProduceHandler(c *gin.Context) {
	params := c.Request.URL.Query()
	targetCodes := params["Produce Code"]

	// if the internal slice is empty, no point in continuing
	if len(store.ProduceItems) == 0 {
		c.JSON(http.StatusOK, ResultSet{Status: "There are currently no produce items.", Results: nil})
		return
	}

	// make a channel big enough to hold the current list of produce items
	produceList := make(chan []Produce, len(store.ProduceItems))
	preLength := len(store.ProduceItems) // get length of slice before changes, used for a check later on
	var notFound []string
	for _, code := range targetCodes {
		// confirm the passed in code is of a valid format
		// make a temporary Produce struct to pass into IsValid
		err := IsValid(Produce{Name: "Test", ProduceCode: code, Price: 1.00})
		if err != nil {
			c.JSON(http.StatusBadRequest, "the inputted codes is of an invalid format")
			return
		}

		go store.RemoveProduceChannel(code, produceList) // delete one item

		select {
		case newList := <-produceList: // if channel receives data, update the internal DB
			store.ProduceItems = newList
		case <-time.After(1 * time.Second): // item not found so the channel is not updated
			notFound = append(notFound, code)
		}
	}

	close(produceList)

	fmt.Println(preLength, len(store.ProduceItems))
	// check if the internal slice was adjusted at all
	if preLength == len(store.ProduceItems) {
		c.JSON(http.StatusBadRequest, ResultSet{Status: "No items were deleted, please make sure your code(s) are inputted correctly.", Results: nil})
	} else {
		// default status message, assuming notFound is empty
		status := "Delete successful, results shows the current state of the DB."
		if len(notFound) > 0 {
			status = fmt.Sprintf("While some deletions were successful, some were not. Note this list of items that were not found: %v", notFound)
		}
		// return whatever status is current and the state of the ProduceItems slice
		c.JSON(http.StatusOK, ResultSet{Status: status, Results: store.ProduceItems})
	}
}

// display a message when user accesses the root page, telling them to refer to the README for info.
func rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "Please refer to the README for information on the current endpoints.")
}

// APIMain acts as the root for the REST API.
func APIMain(store *Store) {
	router := gin.Default()
	router.Use(cors.Default())

	/* decided to use the /produce/ prefix to each API endpoint, even if it isn't strictly necessary.
	my thinking is that in the future that if this is built for production use, we may want another set of endpoints
	for manipulating a store e.g /store/.
	*/
	// API GET endpoints
	router.GET("/", rootHandler)
	router.GET("/produce/getall", store.getAllHandler)
	router.GET("/produce/getitem", store.getProduceHandler)

	// API POST endpoints
	router.POST("/produce/add", store.addProduceHandler)
	router.DELETE("/produce/delete", store.deleteProduceHandler)

	fmt.Println("API has started, please refer to the README for information on the current endpoints.")

	// create port variable that initially assumes default 8080 port, but changes if the PORT env variable is setup
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	fmt.Printf("API is running at: localhost:%v\n", port)

	router.Run()
}
