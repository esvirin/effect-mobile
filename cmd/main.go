package main

import (
	"effect-mobile/internal/db"
	"effect-mobile/internal/handlers"
	"effect-mobile/internal/repository"

	_ "effect-mobile/docs"

	"github.com/gin-gonic/gin"
)

// @title Effect Mobile API
// @version 1.0
// @description API для управления подписками
// @host localhost:8080
// @BasePath /

func main() {
	database := db.InitDB()
	repo := repository.NewSubscriptionRepository(database)
	handler := handlers.NewSubscriptionHandler(repo)

	r := gin.Default()

	// @Summary Создать подписку
	// @Description Создает новую подписку
	// @Tags subscriptions
	// @Accept json
	// @Produce json
	// @Param subscription body handlers.SubscriptionRequest true "Данные подписки"
	// @Success 201 {object} handlers.SubscriptionResponse
	// @Router /subscriptions [post]
	r.POST("/subscriptions", handler.Create)

	// @Summary Получить подписку по ID
	// @Description Возвращает подписку по её ID
	// @Tags subscriptions
	// @Produce json
	// @Param id path int true "ID подписки"
	// @Success 200 {object} handlers.SubscriptionResponse
	// @Router /subscriptions/{id} [get]
	r.GET("/subscriptions/:id", handler.Get)

	// @Summary Список подписок
	// @Description Возвращает все подписки
	// @Tags subscriptions
	// @Produce json
	// @Success 200 {array} handlers.SubscriptionResponse
	// @Router /subscriptions [get]
	r.GET("/subscriptions", handler.List)

	// @Summary Обновить подписку
	// @Description Обновляет данные подписки по ID
	// @Tags subscriptions
	// @Accept json
	// @Produce json
	// @Param id path int true "ID подписки"
	// @Param subscription body handlers.SubscriptionRequest true "Данные подписки"
	// @Success 200 {object} handlers.SubscriptionResponse
	// @Router /subscriptions/{id} [put]
	r.PUT("/subscriptions/:id", handler.Update)

	// @Summary Удалить подписку
	// @Description Удаляет подписку по ID
	// @Tags subscriptions
	// @Param id path int true "ID подписки"
	// @Success 204 {string} string "deleted"
	// @Router /subscriptions/{id} [delete]
	r.DELETE("/subscriptions/:id", handler.Delete)

	// @Summary Сумма подписок
	// @Description Возвращает суммарное значение подписок
	// @Tags subscriptions
	// @Produce json
	// @Success 200 {object} map[string]float64
	// @Router /subscriptions/sum [get]
	r.GET("/subscriptions/sum", handler.Sum)

	r.Run(":8080")
}
