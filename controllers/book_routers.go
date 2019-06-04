package controllers

import (
	"github.com/gin-gonic/gin"
)

func BookRouter(r *gin.Engine) {
	book := r.Group("/api/v1/book")
	book.GET("/search", SearchBooks)
}
