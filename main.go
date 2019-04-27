package main

import (
	"github.com/gin-gonic/gin"
	"seu-job-mina-service/db"
	"seu-job-mina-service/middleware"
)

const Port = "3001"
func main() {

	db.Connect()

	r := gin.Default()

	r.Use(middleware.ErrorWrapper)
	r.Use(middleware.MongoConnect)
	r.Use(middleware.Cors)

	r.Run(":" + Port)
}
