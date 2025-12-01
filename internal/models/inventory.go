package models

import (
	"github.com/google/uuid"
	"time"
)

type Inventory struct {
	Model
	ProductID              uuid.UUID  `json:"product_id" db:"product_id"`
	BatchID                string     `json:"batch_id" db:"batch_id"`
	Quantity               float64    `json:"quantity" db:"quantity"`
	ReceivedDate           time.Time  `json:"received_date" db:"received_date"`
	ReceivedFromSupplierID *string    `json:"received_from_supplier_id" db:"received_from_supplier_id"`
	StorageLocation        string     `json:"storage_location" db:"storage_location"`
	ExpiryDate             time.Time  `json:"expiry_date" db:"expiry_date"`
	FrozenDate             *time.Time `json:"frozen_date" db:"frozen_date"`
	Status                 string     `json:"status" db:"status"`
	Notes                  string     `json:"notes" db:"notes"`

	// Computed Fields
	DaysUntilExpiry int    `json:"days_until_expiry"`
	AlertLevel      string `json:"alert_level"`
	ProductName     string `json:"product_name"`
}

type CreateInventoryRequest struct {
	ProductID              uuid.UUID  `json:"product_id" binding:"required,uuid"`
	BatchID                string     `json:"batch_id" binding:"required"`
	Quantity               float64    `json:"quantity" binding:"required"`
	UnitOfMeasure          string     `json:"unit_of_measure"` // From Product
	ReceivedDate           time.Time  `json:"received_date" binding:"required"`
	ReceivedFromSupplierID *string    `json:"received_from_supplier_id"`
	StorageLocation        string     `json:"storage_location" binding:"required"`
	FrozenDate             *time.Time `json:"frozen_date"`
	Notes                  string     `json:"notes"`
}

type UpdateInventoryRequest struct {
	Quantity        float64 `json:"quantity"`
	UnitOfMeasure   string  `json:"unit_of_measure"` // From Product
	StorageLocation string  `json:"storage_location"`
	Status          string  `json:"status" binding:"omitempty,oneof=available reserved sold expired"`
	Notes           string  `json:"notes"`
}

type InventoryFilter struct {
	ProductID       uuid.UUID
	Code            string
	Status          string
	StorageLocation string
	AlertLevel      string
	ExpiryDateFrom  *time.Time
	ExpiryDateTo    *time.Time
	Page            int
	Limit           int
	SortBy          string
	Order           string
}
