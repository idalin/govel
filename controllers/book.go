package controllers

import (
	"govel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchBooks(c *gin.Context) {
	key := c.Query("q")
	if key == "" {
		c.JSON(http.StatusExpectationFailed, Response{http.StatusExpectationFailed, "no key found.", nil})
		return
	}

	c.JSON(http.StatusOK, Response{http.StatusOK, "success", models.SearchBooks(key)})
}
