package dto

type CreateSupplierRequest struct {
	Name    string `json:"name" binding:"required "`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Notes   string `json:"notes"`
}

type UpdateSupplierRequest struct {
	Name    string `json:"name" binding:"required "`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Notes   string `json:"notes"`
}
