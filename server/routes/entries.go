package routes

import (
	""github.com/gin-gonic/gin""
)

func AddEntry(c *gin.Context) {

}

func GetEntries(c *gin.Context) {

}

func GetEntryById(c *gin.Context) {

}

func GetEntryByIngredient(c *gin.Context) {

}

func UpdateEntry(c *gin.Context) {

}

func UpdateIngredient(c *gin.Context) {

}

func DeleteEntry(c *gin.Context) {
	entryID := c.params.ByName("id")
}
