package flashsaleservice

import (
	"context"
	"flashsale-service/internal/service"
	"flashsale-service/protos/flashSale_proto"
	"fmt"
)

type FlashSaleGrpc struct {
	flashSale_proto.UnimplementedFlashSaleServiceServer
	s service.FlashSaleService
}

func NewFlashSaleGrpc(s service.FlashSaleService) *FlashSaleGrpc {
	return &FlashSaleGrpc{s: s}
}

func (f *FlashSaleGrpc) CreateFlashSale(ctx context.Context, req *flashSale_proto.CreateFlashSaleRequest) (*flashSale_proto.CreateFlashSaleResponse, error){
	res, err := f.s.CreateflashSale(req)
	if err != nil {
		return nil, fmt.Errorf("create flash sale error: %v", err)
	}

	return res, nil
}

func (f *FlashSaleGrpc) GetActiveFlashSales(ctx context.Context, req *flashSale_proto.GetActiveFlashSalesRequest) (*flashSale_proto.GetActiveFlashSalesResponse, error){
	res, err := f.s.GetActiveflashSales(req)
	if err != nil {
		return nil, fmt.Errorf("get active flash sale error: %v", err)
	}

	return res, nil
}

func (f *FlashSaleGrpc) UpdateFlashSale(ctx context.Context, req *flashSale_proto.UpdateFlashSaleRequest) (*flashSale_proto.UpdateFlashSaleResponse, error){
	res, err := f.s.UpdateflashSale(req)
	if err != nil {
		return nil, fmt.Errorf("update flash sale error: %v", err)
	}

	return res, nil
}

func (f *FlashSaleGrpc) DeleteFlashSale(ctx context.Context,req *flashSale_proto.DeleteFlashSaleRequest) (*flashSale_proto.DeleteFlashSaleResponse, error){
	res, err := f.s.DeleteflashSale(req)
	if err != nil {
		return nil, fmt.Errorf("delete flash sale error: %v", err)
	}

	return res, nil
}

func (f *FlashSaleGrpc) PurchaseProduct(ctx context.Context, req *flashSale_proto.PurchaseProductRequest) (*flashSale_proto.PurchaseProductResponse, error){
	res, err := f.s.Purchaseproduct(req)
	if err != nil {
		return nil, fmt.Errorf("purchase product error: %v", err)
	}

	return res, nil
}