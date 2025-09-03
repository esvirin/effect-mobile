package repository

import (
	"time"

	"effect-mobile/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(sub *models.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *SubscriptionRepository) GetByID(id uuid.UUID) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.First(&sub, "id = ?", id).Error
	return &sub, err
}

func (r *SubscriptionRepository) List() ([]models.Subscription, error) {
	var subs []models.Subscription
	err := r.db.Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepository) Update(sub *models.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *SubscriptionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Subscription{}, "id = ?", id).Error
}

func (r *SubscriptionRepository) Sum(userID *uuid.UUID, serviceName *string, from time.Time, to time.Time) (int, error) {
	var sum int
	query := r.db.Model(&models.Subscription{}).
		Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", to, from)

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if serviceName != nil {
		query = query.Where("service_name = ?", *serviceName)
	}

	err := query.Select("COALESCE(SUM(price),0)").Scan(&sum).Error
	return sum, err
}
