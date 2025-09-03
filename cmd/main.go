package main

import (
	"effect-mobile/internal/db"
	"effect-mobile/internal/handlers"
	"effect-mobile/internal/repository"

	"github.com/gin-gonic/gin"
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
