package models

import "time"

type PricingHistory struct {
	ProductID         string     `json:"product_id" db:"product_id"`
	BuyingPricePerKg  float64    `json:"buying_price_per_kg" db:"buying_price_per_kg"`
	SellingPricePerKg float64    `json:"selling_price_per_kg" db:"selling_price_per_kg"`
	MarginPercentage  float64    `json:"margin_percentage" db:"margin_percentage"`
	Currency          string     `json:"currency" db:"currency"`
	EffectiveFrom     time.Time  `json:"effective_from" db:"effective_from"`
	EffectiveUntil    *time.Time `json:"effective_until" db:"effective_until"`
	Reason            string     `json:"reason" db:"reason"`
	CreatedBy         *string    `json:"created_by" db:"created_by"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	IsActive          bool       `json:"is_active"`
	ProductName       string     `json:"product_name"`
}

type CurrentPrice struct {
	ProductID         string    `json:"product_id" db:"product_id"`
	ProductName       string    `json:"product_name" db:"product_name"`
	BuyingPricePerKg  float64   `json:"buying_price_per_kg" db:"buying_price_per_kg"`
	SellingPricePerKg float64   `json:"selling_price_per_kg" db:"selling_price_per_kg"`
	MarginPercentage  float64   `json:"margin_percentage" db:"margin_percentage"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type UpdatePriceRequest struct {
	ProductID         string  `json:"product_id" binding:"required,uuid"`
	BuyingPricePerKg  float64 `json:"buying_price_per_kg" binding:"required,gt=0"`
	SellingPricePerKg float64 `json:"selling_price_per_kg" binding:"required,gt=0"`
	Currency          string  `json:"currency" binding:"omitempty,len=3"`
	Reason            string  `json:"reason" binding:"required"`
}
