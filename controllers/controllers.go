package controllers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func StartGin(port string, r *gin.Engine) {
	BookRouter(r)
	WSRouter(r)
	r.Run(port)
}
