package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rccInventory/internal/models"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
	GetAll(ctx context.Context, page, limit int) ([]models.Product, int64, error)
	Update(ctx context.Context, id uuid.UUID, updates *models.UpdateProductRequest) error
	Delete(ctx context.Context, id uuid.UUID) error

	// GetBySupplierID(ctx context.Context, supplierID string, page, limit int) ([]*models.Product, int64, error)

}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(ctx context.Context, product *models.Product) error {
	result := r.db.WithContext(ctx).Create(product)
	return result.Error
}

func (r *productRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	var product models.Product
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}

	return &product, result.Error
}

func (r *productRepo) GetAll(ctx context.Context, page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	result := r.db.WithContext(ctx).
		Where("status = ?", "active").
		Offset(offset).
		Limit(limit).
		Find(&products)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Get total count
	r.db.WithContext(ctx).
		Where("status = ?", "active").
		Model(&models.Product{}).
		Count(&total)

	return products, total, nil

}

func (r *productRepo) Update(ctx context.Context, id uuid.UUID, updates *models.UpdateProductRequest) error {
	return r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error
}

func (r *productRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Update("status", "discontinued").Error
}
