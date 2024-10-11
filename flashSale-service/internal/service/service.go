package service

import (
	productclients "flashsale-service/internal/clients/product_clients"
	flashsale "flashsale-service/internal/entity/flashSale"
	"flashsale-service/internal/infrastructura/repository"
	"flashsale-service/protos/flashSale_proto"
	"fmt"
	"log"
)

type FlashSaleService struct {
	repo repository.FlashSaleRepository
}

func NewFlashSaleService(repo repository.FlashSaleRepository) *FlashSaleService {
	return &FlashSaleService{repo: repo}
}

func (f *FlashSaleService) CreateflashSale(req *flashSale_proto.CreateFlashSaleRequest) (*flashSale_proto.CreateFlashSaleResponse, error) {
	var products []flashsale.FlashSaleProduct
	for _, product := range req.Products {
		res, err := productclients.GetProductbyId(product.ProductId)
		if err != nil {
			log.Println(product.ProductId, "product not found")
			continue
		}
		err = productclients.UpdateProductDiscountPrice(product.ProductId, product.DescountPrice)
		if err != nil {
			log.Println("update product error: ", err)
			continue
		}
		products = append(products, flashsale.FlashSaleProduct{
			ProductID:     res.ProductId,
			Name:          res.Name,
			Description:   res.Description,
			OriginalPrice: res.Price,
			DiscountPrice: product.DescountPrice,
			Stock:         res.StockQuantity,
			IsActive:      true,
		})
	}
	res, err := f.repo.InsertMongoDb(flashsale.CreateFlashSaleRequest{
		Name:        req.Name,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Products:    products,
	})
	if err != nil {
		log.Printf("Error inserting flash sale into MongoDB: %v", err)
		return nil, fmt.Errorf("create mongodb error: %w", err)
	}

	protoEvent := convertEventToProto(res.Event)

	return &flashSale_proto.CreateFlashSaleResponse{
		Event: protoEvent,
	}, nil
}

func (f *FlashSaleService) UpdateflashSale(req *flashSale_proto.UpdateFlashSaleRequest) (*flashSale_proto.UpdateFlashSaleResponse, error) {
	res, err := f.repo.UpdateFlashsaleMongodb(flashsale.UpdateFlashSaleRequest{
		FlashSaleId: req.Id,
		Name:        req.Name,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Products:    convertProtoProductsToDomain(req.Products),
	})
	if err != nil {
		log.Println("Error  updated flash sale into mongodb:", err)
		return nil, fmt.Errorf("error updated flash sale into mongodb: %v", err)
	}

	return &flashSale_proto.UpdateFlashSaleResponse{
		Event: convertEventToProto(res.Event),
	}, nil
}

func (f *FlashSaleService) DeleteflashSale(req *flashSale_proto.DeleteFlashSaleRequest) (*flashSale_proto.DeleteFlashSaleResponse, error) {
	_, err := f.repo.DeleteFlashsaleMongodb(flashsale.DeleteFlashSaleRequest{FlashSaleId: req.Id})
	if err != nil {
		log.Println("error deleted flash sale into mongodb: ", err)
		return nil, fmt.Errorf("error deleted flash sale into mongodb: %v", err)
	}

	return &flashSale_proto.DeleteFlashSaleResponse{Message: "flash sale deleted"}, nil
}

func (f *FlashSaleService) ListflashSale(req *flashSale_proto.ListFlashSalesRequest) (*flashSale_proto.ListFlashSalesResponse, error) {
	res, err := f.repo.ListFlashsaleMongodb(flashsale.ListFlashSalesRequest{})
	if err != nil {
		log.Println("error get all flash sale into mongodb: ", err)
		return nil, fmt.Errorf("error get all flash sale into mongodb: %v", err)
	}

	return &flashSale_proto.ListFlashSalesResponse{
		Events: convertEventsToProto(res.Events),
	}, nil
}

func (f *FlashSaleService) GetActiveflashSales(req *flashSale_proto.GetActiveFlashSalesRequest) (*flashSale_proto.GetActiveFlashSalesResponse, error) {
	res, err := f.repo.GetActiveFlashSalesMongodb(flashsale.GetActiveFlashSalesRequest{})
	if err != nil {
		log.Println("error get active flash sales into mongodb: ", err)
		return nil, fmt.Errorf("error get active flash sales into mongodb: %v", err)
	}

	return &flashSale_proto.GetActiveFlashSalesResponse{
		ActiveEvents: convertEventsToProto(res.ActiveEvents),
	}, nil
}

func (f *FlashSaleService) Purchaseproduct(req *flashSale_proto.PurchaseProductRequest) (*flashSale_proto.PurchaseProductResponse, error) {
	res, err := f.repo.PurchaseProductMongodb(flashsale.PurchaseProductRequest{
		ProductID: req.ProductId,
		UserID:    req.UserId,
		Quantity:  req.Quantity,
	})

	if err != nil {
		log.Println("error get purchase product into mongodb: ", err)
		return nil, fmt.Errorf("error get purchase product into mongodb: %v", err)
	}

	return &flashSale_proto.PurchaseProductResponse{
		Message: "Prodcuts",
		Product: convertProductToProto(res.Product),
	}, nil
}

func convertProductToProto(product flashsale.FlashSaleProduct) *flashSale_proto.FlashSaleProduct {
	return &flashSale_proto.FlashSaleProduct{
		ProductId:     product.ProductID,
		Name:          product.Name,
		Description:   product.Description,
		OriginalPrice: product.OriginalPrice,
		DiscountPrice: product.DiscountPrice,
		Stock:         product.Stock,
		IsActive:      product.IsActive,
	}
}

func convertProtoProductsToDomain(protoProducts []*flashSale_proto.FlashSaleProduct) []flashsale.FlashSaleProduct {
	var products []flashsale.FlashSaleProduct
	for _, product := range protoProducts {
		res, err := productclients.GetProductbyId(product.ProductId)
		if err != nil {
			log.Println(product.ProductId, "product not found")
			continue
		}
		err = productclients.UpdateProductDiscountPrice(product.ProductId, res.DiscountPrice)
		if err != nil {
			log.Println("update product error: ", err)
			continue
		}
		products = append(products, flashsale.FlashSaleProduct{
			ProductID:     res.ProductId,
			Name:          res.Name,
			Description:   res.Description,
			OriginalPrice: res.Price,
			DiscountPrice: res.DiscountPrice,
			Stock:         res.StockQuantity,
			IsActive:      true,
		})
	}
	return products
}

func convertEventsToProto(events []flashsale.FlashSaleEvent) []*flashSale_proto.FlashSaleEvent {
	var protoEvents []*flashSale_proto.FlashSaleEvent
	for _, event := range events {
		protoEvents = append(protoEvents, convertEventToProto(event))
	}
	return protoEvents
}

func convertEventToProto(event flashsale.FlashSaleEvent) *flashSale_proto.FlashSaleEvent {
	return &flashSale_proto.FlashSaleEvent{
		Id:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Products:    convertProductsToProto(event.Products),
		IsActive:    event.IsActive,
	}
}

func convertProductsToProto(domainProducts []flashsale.FlashSaleProduct) []*flashSale_proto.FlashSaleProduct {
	var products []*flashSale_proto.FlashSaleProduct
	for _, product := range domainProducts {
		products = append(products, &flashSale_proto.FlashSaleProduct{
			ProductId:     product.ProductID,
			Name:          product.Name,
			Description:   product.Description,
			OriginalPrice: product.OriginalPrice,
			DiscountPrice: product.DiscountPrice,
			Stock:         product.Stock,
			IsActive:      product.IsActive,
		})
	}
	return products
}
