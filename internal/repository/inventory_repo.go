package repository

import (
	"context"
	"github.com/google/uuid"
	"rccInventory/internal/models"
)

type InventoryRepository interface {
	Create(ctx context.Context, inventory *models.Inventory) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Inventory, error)
	GetByBatchID(ctx context.Context, batchID string) (*models.Inventory, error)
	GetFiltered(ctx context.Context, filter *models.InventoryFilter) ([]models.Inventory, int64, error)
	Update(ctx context.Context, id uuid.UUID, updates *models.UpdateInventoryRequest) error
	GetExpiringItems(ctx context.Context, daysThreshold int) ([]models.Inventory, error)
	GetByProductID(ctx context.Context, productID uuid.UUID) (*models.Inventory, error)
}
