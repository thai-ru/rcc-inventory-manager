package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rccInventory/internal/models"
)

type InventoryRepository interface {
	Create(ctx context.Context, inventory *models.Inventory) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Inventory, error)
	GetByBatchID(ctx context.Context, batchID string) (*models.Inventory, error)
	GetFiltered(ctx context.Context, filter *models.InventoryFilter) ([]models.Inventory, int64, error)
	Update(ctx context.Context, id uuid.UUID, updates *models.UpdateInventoryRequest) error
	GetExpiringItems(ctx context.Context, daysThreshold int) ([]models.Inventory, error)
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]models.Inventory, error)
	UpdateQuantity(ctx context.Context, id uuid.UUID, newQuantity float64) error
}

type inventoryRepo struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepo{db: db}
}

func (r *inventoryRepo) Create(ctx context.Context, inventory *models.Inventory) error {
	result := r.db.WithContext(ctx).Create(inventory)
	return result.Error
}

func (r *inventoryRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Inventory, error) {
	var inventory models.Inventory
	result := r.db.WithContext(ctx).
		Select("i.*, p.name as product_name").
		Table("inventory i").
		Joins("LEFT JOIN products p ON i.product_id = p.id").
		Where("i.id = ?", id).
		Scan(&inventory)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("inventory not found")
	}

	return &inventory, result.Error
}

func (r *inventoryRepo) GetByBatchID(ctx context.Context, batchID string) (*models.Inventory, error) {
	var inventory models.Inventory

	result := r.db.WithContext(ctx).Where("batch_id = ?", batchID).First(&inventory)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("batch not found")
	}

	return &inventory, result.Error
}

func (r *inventoryRepo) GetFiltered(ctx context.Context, filter *models.InventoryFilter) ([]models.Inventory, int64, error) {
	var inventories []models.Inventory
	var total int64

	query := r.db.WithContext(ctx).
		Select("i.*, p.name as product_name").
		Table("inventory i").
		Joins("LEFT JOIN products p ON i.product_id = p.id")

	// Apply filters
	if filter.ProductID != uuid.Nil {
		query = query.Where("i.product_id = ?", filter.ProductID)
	}
	if filter.Status != "" {
		query = query.Where("i.status = ?", filter.Status)
	}
	if filter.StorageLocation != "" {
		query = query.Where("i.storage_location = ?", filter.StorageLocation)
	}
	if filter.ExpiryDateFrom != nil {
		query = query.Where("i.expiry_date >= ?", filter.ExpiryDateFrom)
	}
	if filter.ExpiryDateTo != nil {
		query = query.Where("i.expiry_date <= ?", filter.ExpiryDateTo)
	}

	// Count total before pagination
	query.Model(&models.Inventory{}).Count(&total)

	// Apply sorting
	sortBy := "i.created_at"
	if filter.SortBy != "" {
		sortBy = "i." + filter.SortBy
	}
	order := "DESC"
	if filter.Order == "asc" {
		order = "ASC"
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.Limit
	result := query.
		Order(sortBy + " " + order).
		Offset(offset).
		Limit(filter.Limit).
		Scan(&inventories)

	return inventories, total, result.Error
}

func (r *inventoryRepo) Update(ctx context.Context, id uuid.UUID, updates *models.UpdateInventoryRequest) error {
	return r.db.WithContext(ctx).Model(&models.Inventory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *inventoryRepo) GetExpiringItems(ctx context.Context, daysThreshold int) ([]models.Inventory, error) {
	var inventories []models.Inventory
	result := r.db.WithContext(ctx).
		Where("status = ? AND expiry_date <= NOW() + INTERNAL '? days'", "available", daysThreshold).
		Find(&inventories)

	return inventories, result.Error
}

func (r *inventoryRepo) GetByProductID(ctx context.Context, productID uuid.UUID) ([]models.Inventory, error) {
	var inventories []models.Inventory
	result := r.db.WithContext(ctx).
		Where("product_id = ? AND status", productID, "available").
		Find(&inventories)

	return inventories, result.Error
}

func (r *inventoryRepo) UpdateQuantity(ctx context.Context, id uuid.UUID, newQuantity float64) error {
	return r.db.WithContext(ctx).Model(&models.Inventory{}).Where("id = ?", id).Update("quantity", newQuantity).Error
}
