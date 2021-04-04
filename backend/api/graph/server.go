package main

import (
	"os"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	middleware.InitDB()
	port := lib.GetDefVal("PORT", "8080")

	r.Static("/documents/", "./documents")

	r.Use(middleware.SetContext())
	r.Use(middleware.CORSMiddleware())

	r.POST("/query", middleware.GraphqlHandler())
	r.GET("/", middleware.PlaygroundHandler())
	lib.LogInfo("main", "🚀 Server listening on :"+port)
	if err := r.Run(); err != nil {
		lib.LogError("main", err.Error())
		os.Exit(1)
	}
}
