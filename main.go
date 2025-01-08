package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

//item represents data about a item in a grocery store
type item struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Type string  `json:"type"`
    Price  float64 `json:"price"`
}

// items slice containing items available in a grocery store
var items = []item{
    {ID: "1", Name: "Banana", Type: "Fruit", Price: 40},
    {ID: "2", Name: "Tomoto", Type: "Vegetable", Price: 30},
    {ID: "3", Name: "Cocumber", Type: "Vegetable", Price: 20},
}

func main() {
    router := gin.Default()
    router.GET("/items", getItems)
    router.GET("/items/:id", getItemByID)
    router.POST("/items", addItem)
	router.DELETE("/items/:id",deleteItem)
	router.PUT("/items/:id",updateItem)
    router.Run(":8080")
}

// getItems responds with the list of all getItems present in a grocery store as JSON.
func getItems(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, items)
}

// getItemByID searches the item with the same id in grocery store and sends response as JSON.
func getItemByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of items in grocery store, looking for
    // an item whose ID value matches the parameter.
    for _, a := range items {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found in grocery store."})
}

// adddItem adds the item from JSON received into grocery store inventory.
func addItem(c *gin.Context) {
    var newItem item

    // Call BindJSON to bind the received JSON to
    // item.
    if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "Bad Request ": err.Error(),
        })
        return
    }

    // Add the new item to the grocery store.
    items = append(items, newItem)
    c.IndentedJSON(http.StatusCreated, newItem)
}

//deleteItem delets the item from the grocery store .
func deleteItem(c *gin.Context) {
	id := c.Param("id")
	// Loop through the list of items in grocery store, looking for
    // an item whose ID value matches the parameter and delete it
	for i, a := range items {
        if a.ID == id {
            items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found in grocery store."})
}

//updateItem searches the item with the same id in grocery store and updates it.
func updateItem(c *gin.Context) {
	id := c.Param("id")
	var newItem item
	
	// Call BindJSON to bind the received JSON to
    // item.
    if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "Bad Request ": err.Error(),
        })
        return
    }

	for i, a := range items {
		if a.ID == id {
			newItem.ID = id // set the id of newItem
			items[i] = newItem
			c.JSON(http.StatusOK, newItem)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found in grocery store."})
}