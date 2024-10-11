package repository

import flashsale "flashsale-service/internal/entity/flashSale"

type FlashSaleRepository interface {
	InsertMongoDb(req flashsale.CreateFlashSaleRequest) (*flashsale.CreateFlashSaleResponse, error)
	UpdateFlashsaleMongodb(req flashsale.UpdateFlashSaleRequest) (*flashsale.UpdateFlashSaleResponse, error)
	DeleteFlashsaleMongodb(req flashsale.DeleteFlashSaleRequest) (*flashsale.DeleteFlashSaleResponse, error)
	ListFlashsaleMongodb(req flashsale.ListFlashSalesRequest) (*flashsale.ListFlashSalesResponse, error)
	GetActiveFlashSalesMongodb(req flashsale.GetActiveFlashSalesRequest) (*flashsale.GetActiveFlashSalesResponse, error)
	PurchaseProductMongodb(req flashsale.PurchaseProductRequest) (*flashsale.PurchaseProductResponse, error)
}
