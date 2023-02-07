package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cache = map[string]string{}

func healthCheck(c *gin.Context) {

	data := map[string]interface{}{

		"status": "ok",
	}

	c.JSON(http.StatusOK, data)

}

func createRecord(c *gin.Context) {

	var record Record
	if c.ShouldBind(&record) == nil {

		cache[record.Key] = record.Value

		data := map[string]interface{}{
			"key":    record.Key,
			"value":  record.Value,
			"status": "created",
		}

		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(400, http.StatusBadRequest)
	}

}

func getRecord(c *gin.Context) {

	var query Query
	if c.ShouldBind(&query) == nil {

		result := cache[query.Key]

		if len(result) > 0 {
			data := map[string]interface{}{
				"key":   query.Key,
				"value": result,
			}

			c.JSON(http.StatusOK, data)

		} else {

			c.JSON(404, http.StatusNotFound)
		}

	}
}

func deleteRecord(c *gin.Context) {

	var query Query
	if c.ShouldBind(&query) == nil {

		delete(cache, query.Key)

		data := map[string]interface{}{
			"status": "deleted",
		}

		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(400, http.StatusBadRequest)
	}
}

func flushKeys(c *gin.Context) {

	cache = map[string]string{}

	data := map[string]interface{}{

		"status": "flushed all keys",
	}

	c.JSON(http.StatusOK, data)

}
