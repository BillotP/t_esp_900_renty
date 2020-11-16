package main

import (
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/lib/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	middleware.InitDB()

	r.Use(middleware.SetContext())
	r.Use(middleware.CORSMiddleware())

	r.POST("/query", middleware.GraphqlHandler())
	r.GET("/", middleware.PlaygroundHandler())
	r.Run()
}
