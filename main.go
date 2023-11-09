package main

import (
	"apirediscache/controllers"
	"apirediscache/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()
	r.GET("/product", controllers.GetAllProducts)
	r.POST("/product", controllers.CreateProduct)
	r.Run()
}
