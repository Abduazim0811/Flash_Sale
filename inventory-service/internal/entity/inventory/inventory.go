package inventory

type CreateInventory struct{
	InventoryID 		string			`json:"order_id"`
	ProductId		string			`json:"product_id"`
	Quantity		int32			`json:"quantity"`
	CreatedAt		string			`json:"created_at"`
	UpdatedAt		string			`json:"updated_at"`
}

type GetbyIdInventoryReq struct {
	ProductId string `json:"product_id"`
}

type GetAllInventoriesReq struct{}

type GetAllInventoriesRes struct {
	Inventories []Inventory `json:"inventories"`
}

type UpdateInventoryRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type UpdateInventoryResponse struct {
	Inventory Inventory `json:"inventory"`
}

type Inventory struct {
	ProductID  string `json:"product_id"`
	Quantity   int32  `json:"quantity"`
	IsActive   bool   `json:"is_active"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

