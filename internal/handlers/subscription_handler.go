package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/subscription-service/internal/models"
	"github.com/subscription-service/internal/repository"
)

type SubscriptionHandler struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionHandler(repo *repository.SubscriptionRepository) *SubscriptionHandler {
	return &SubscriptionHandler{repo: repo}
}

// @Summary Create subscription
// @Accept json
// @Produce json
// @Param data body models.Subscription true "Subscription"
// @Success 201 {object} models.Subscription
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub models.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create subscription"})
		return
	}
	c.JSON(http.StatusCreated, sub)
}

func (h *SubscriptionHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	sub, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func (h *SubscriptionHandler) List(c *gin.Context) {
	subs, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot list"})
		return
	}
	c.JSON(http.StatusOK, subs)
}

func (h *SubscriptionHandler) Update(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var sub models.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub.ID = id
	if err := h.repo.Update(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update"})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *SubscriptionHandler) Sum(c *gin.Context) {
	var (
		userID      *uuid.UUID
		serviceName *string
	)

	if uid := c.Query("user_id"); uid != "" {
		u, err := uuid.Parse(uid)
		if err == nil {
			userID = &u
		}
	}
	if sn := c.Query("service_name"); sn != "" {
		serviceName = &sn
	}

	fromStr := c.Query("from") // "2025-07-01"
	toStr := c.Query("to")
	from, _ := time.Parse("2006-01-02", fromStr)
	to, _ := time.Parse("2006-01-02", toStr)

	sum, err := h.repo.Sum(userID, serviceName, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot calc sum"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
