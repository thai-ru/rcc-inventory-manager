package models

type Product struct {
	Model
	Name          string  `json:"name" db:"name"`
	CutType       string  `json:"cut_type" db:"cut_type"`
	Description   string  `json:"description" db:"description"`
	UnitOfMeasure string  `json:"unit_of_measure" db:"unit_of_measure"`
	SupplierID    *string `json:"supplier_id" db:"supplier_id"`
	Status        string  `json:"status" db:"status"`
}

type CreateProductRequest struct {
	Name          string  `json:"name" binding:"required,min=2"`
	CutType       string  `json:"cut_type" binding:"required"`
	Description   string  `json:"description"`
	UnitOfMeasure string  `json:"unit_of_measure" binding:"required"`
	SupplierID    *string `json:"supplier_id"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"omitempty,oneof=active discontinued"`
}
