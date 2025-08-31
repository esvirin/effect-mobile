package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subscription-service/internal/db"
	"github.com/subscription-service/internal/handlers"
	"github.com/subscription-service/internal/repository"
)

func main() {
	database := db.InitDB()
	repo := repository.NewSubscriptionRepository(database)
	handler := handlers.NewSubscriptionHandler(repo)

	r := gin.Default()

	r.POST("/subscriptions", handler.Create)
	r.GET("/subscriptions/:id", handler.Get)
	r.GET("/subscriptions", handler.List)
	r.PUT("/subscriptions/:id", handler.Update)
	r.DELETE("/subscriptions/:id", handler.Delete)

	r.GET("/subscriptions/sum", handler.Sum)

	r.Run(":8080")
}
