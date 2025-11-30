package models

import "time"

type StockMovement struct {
	ID            string    `json:"id" db:"id"`
	InventoryID   string    `json:"inventory_id" db:"inventory_id"`
	MovementType  string    `json:"movement_type" db:"movement_type"`
	Quantity      int       `json:"quantity" db:"quantity"`
	UnitOfMeasure string    `json:"unit_of_measure" db:"unit_of_measure"`
	Reason        string    `json:"reason" db:"reason"`
	ReferenceID   string    `json:"reference_id" db:"reference_id"`
	CreatedBy     *string   `json:"created_by" db:"created_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type CreateMovementRequest struct {
	InventoryID   string `json:"inventory_id" binding:"required,uuid"`
	MovementType  string `json:"movement_type" binding:"required,oneof=in out adjustment"`
	Quantity      int    `json:"quantity" db:"quantity"`
	UnitOfMeasure string `json:"unit_of_measure" db:"unit_of_measure"`
	Reason        string `json:"reason" binding:"required"`
	ReferenceID   string `json:"reference_id"`
}
