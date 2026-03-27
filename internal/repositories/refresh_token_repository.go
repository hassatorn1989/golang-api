package repositories

import (
	"todo-fiber/internal/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(token *models.RefreshTokenModel) error {
	return r.db.Create(token).Error
}

func (r *RefreshTokenRepository) FindByToken(token string) (*models.RefreshTokenModel, error) {
	var refreshToken models.RefreshTokenModel
	err := r.db.Where("token = ?", token).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepository) Revoke(token string) error {
	return r.db.Model(&models.RefreshTokenModel{}).Where("token = ?", token).Update("revoked", true).Error
}

func (r *RefreshTokenRepository) DeleteExpired() error {
	return r.db.Where("expires_at < ?", gorm.Expr("NOW()")).Delete(&models.RefreshTokenModel{}).Error
}
