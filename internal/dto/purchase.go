package dto

import "github.com/shopspring/decimal"

type CreatePurchaseItemRequest struct {
	ProductID uint            `json:"product_id" binding:"required"`
	Quantity  decimal.Decimal `json:"quantity" binding:"required"`
	Price     decimal.Decimal `json:"price" binding:"required"`
}

type CreatePurchaseRequest struct {
	SupplierID uint                        `json:"supplier_id" binding:"required"`
	Notes      string                      `json:"notes"`
	Items      []CreatePurchaseItemRequest `json:"items" binding:"required,min=1"`
}

type UpdatePurchaseStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending received cancelled"`
}
