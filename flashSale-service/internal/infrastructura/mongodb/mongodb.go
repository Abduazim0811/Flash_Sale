package mongodb

import (
	"context"
	flashsale "flashsale-service/internal/entity/flashSale"
	"flashsale-service/internal/infrastructura/repository"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlashSaleMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewFlashSaleMongodb(client *mongo.Client, collection *mongo.Collection) repository.FlashSaleRepository {
	return &FlashSaleMongodb{client: client, collection: collection}
}

func (f *FlashSaleMongodb) InsertMongoDb(req flashsale.CreateFlashSaleRequest) (*flashsale.CreateFlashSaleResponse, error) {
	fmt.Println("salom")
	flashSaleEvent := flashsale.FlashSaleEvent{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Products:    req.Products,
		IsActive:    true,
	}

	_, err := f.collection.InsertOne(context.Background(), flashSaleEvent)
	if err != nil {
		log.Println("failed to insert flash sale event:", err)
		return nil, fmt.Errorf("failed to insert flash sale event: %w", err)
	}

	response := &flashsale.CreateFlashSaleResponse{
		Event: flashSaleEvent,
	}


	log.Println("Kafka message sent successfully")
	return response, nil
}

func (f *FlashSaleMongodb) UpdateFlashsaleMongodb(req flashsale.UpdateFlashSaleRequest) (*flashsale.UpdateFlashSaleResponse, error) {
	filter := bson.M{"id": req.FlashSaleId}

	update := bson.M{}

	if req.Name != "" {
		update["name"] = req.Name
	}
	if req.Description != "" {
		update["description"] = req.Description
	}
	if req.StartTime != "" {
		update["start_time"] = req.StartTime
	}
	if req.EndTime != "" {
		update["end_time"] = req.EndTime
	}
	if len(req.Products) > 0 {
		update["products"] = req.Products
	}

	result, err := f.collection.UpdateOne(context.Background(), filter, bson.M{"$set": update})
	if err != nil {
		log.Println("failed to update flash sale event:", err)
		return nil, fmt.Errorf("failed to update flash sale event: %w", err)
	}

	if result.MatchedCount == 0 {
		log.Printf("no flash sale event found with id: %s\n", req.FlashSaleId)
		return nil, fmt.Errorf("no flash sale event found with id: %s", req.FlashSaleId)
	}

	flashSaleEvent := flashsale.FlashSaleEvent{
		ID:          req.FlashSaleId,
		Name:        req.Name,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Products:    req.Products,
		IsActive:    true,
	}

	response := &flashsale.UpdateFlashSaleResponse{
		Event: flashSaleEvent,
	}

	return response, nil
}

func (f *FlashSaleMongodb) DeleteFlashsaleMongodb(req flashsale.DeleteFlashSaleRequest) (*flashsale.DeleteFlashSaleResponse, error) {
	filter := bson.M{"flashsale_id": req.FlashSaleId}

	result, err := f.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("failed to delete flash sale event:", err)
		return nil, fmt.Errorf("failed to delete flash sale event: %w", err)
	}

	if result.DeletedCount == 0 {
		log.Printf("no flash sale event found with id: %s\n", req.FlashSaleId)
		return nil, fmt.Errorf("no flash sale event found with id: %s", req.FlashSaleId)
	}

	response := &flashsale.DeleteFlashSaleResponse{
		Message: "Flash sale event successfully deleted",
	}

	return response, nil
}

func (f *FlashSaleMongodb) ListFlashsaleMongodb(req flashsale.ListFlashSalesRequest) (*flashsale.ListFlashSalesResponse, error) {
	cursor, err := f.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("failed to retrieve flash sales:", err)
		return nil, fmt.Errorf("failed to retrieve flash sales: %w", err)
	}
	defer cursor.Close(context.Background())

	var events []flashsale.FlashSaleEvent

	for cursor.Next(context.Background()) {
		var event flashsale.FlashSaleEvent
		if err := cursor.Decode(&event); err != nil {
			log.Println("failed to decode flash sale event:", err)
			return nil, fmt.Errorf("failed to decode flash sale event: %w", err)
		}
		events = append(events, event)
	}

	if err := cursor.Err(); err != nil {
		log.Println("error occurred while iterating cursor:", err)
		return nil, fmt.Errorf("error occurred while iterating cursor: %w", err)
	}

	response := &flashsale.ListFlashSalesResponse{
		Events: events,
	}

	return response, nil
}

func (f *FlashSaleMongodb) GetActiveFlashSalesMongodb(req flashsale.GetActiveFlashSalesRequest) (*flashsale.GetActiveFlashSalesResponse, error) {
	filter := bson.M{"isactive": true}

	cursor, err := f.collection.Find(context.Background(), filter)
	if err != nil {
		log.Println("failed to retrieve active flash sales:", err)
		return nil, fmt.Errorf("failed to retrieve active flash sales: %w", err)
	}
	defer cursor.Close(context.Background())

	var activeEvents []flashsale.FlashSaleEvent

	for cursor.Next(context.Background()) {
		var event flashsale.FlashSaleEvent
		if err := cursor.Decode(&event); err != nil {
			log.Println("failed to decode active flash sale event:", err)
			return nil, fmt.Errorf("failed to decode active flash sale event: %w", err)
		}
		activeEvents = append(activeEvents, event)
	}

	if err := cursor.Err(); err != nil {
		log.Println("error occurred while iterating cursor:", err)
		return nil, fmt.Errorf("error occurred while iterating cursor: %w", err)
	}

	response := &flashsale.GetActiveFlashSalesResponse{
		ActiveEvents: activeEvents,
	}

	return response, nil
}

func (f *FlashSaleMongodb) PurchaseProductMongodb(req flashsale.PurchaseProductRequest) (*flashsale.PurchaseProductResponse, error) {
	var product flashsale.FlashSaleProduct
	err := f.collection.FindOne(context.TODO(), bson.M{"products.product_id": req.ProductID}).Decode(&product)
	if err != nil {
		log.Printf("failed to find product with ID %s: %v", req.ProductID, err)
		return nil, fmt.Errorf("failed to find product with ID %s: %w", req.ProductID, err)
	}

	if product.Stock < req.Quantity {
		log.Printf("not enough stock for product ID %s: requested %d, available %d", req.ProductID, req.Quantity, product.Stock)
		return nil, fmt.Errorf("not enough stock for product ID %s: requested %d, available %d", req.ProductID, req.Quantity, product.Stock)
	}

	update := bson.M{"$inc": bson.M{"products.$.stock": -req.Quantity}}
	_, err = f.collection.UpdateOne(context.TODO(), bson.M{"products.product_id": req.ProductID}, update)
	if err != nil {
		log.Printf("failed to update stock for product ID %s: %v", req.ProductID, err)
		return nil, fmt.Errorf("failed to update stock for product ID %s: %w", req.ProductID, err)
	}

	response := &flashsale.PurchaseProductResponse{
		Message: "Product purchased successfully",
		Product: product,
	}

	log.Printf("Product ID %s purchased successfully. Quantity: %d", req.ProductID, req.Quantity)
	return response, nil
}
