package product

type CreateProductReq struct {
    ProductID      string  `json:"product_id"`
    Name           string  `json:"name"`
    Description    string  `json:"description"`
    Price          float32 `json:"price"`
    DiscountPrice  float32 `json:"discount_price"`
    StockQuantity  int32   `json:"stock_quantity"`
}

type CreateProductRes struct {
    ProductID string `json:"product_id"`
    Message   string `json:"message"`
}

type Product struct {
    ProductID      string  `json:"product_id"`
    Name           string  `json:"name"`
    Description    string  `json:"description"`
    Price          float32 `json:"price"`
    DiscountPrice  float32 `json:"discount_price"`
    StockQuantity  int32   `json:"stock_quantity"`
}

type ListProductsRes struct {
    Products []Product `json:"products"`
}

type ListProductsReq struct{}

type GetProductReq struct {
    ProductID string `json:"product_id"`
}

type UpdateProductRes struct {
    Message string `json:"message"`
}

type GetProductOrderReq struct{
    Product_id string   `json:"product_id"`
    StockQuantity   int32    `json:"stock_quantity"`
}

type GetProductOrderres struct{
    TotalAmount  float32    `json:"total_amount"`
}