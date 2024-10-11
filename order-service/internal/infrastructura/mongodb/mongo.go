package mongodb

import (
	"context"
	"fmt"
	"log"

	inventoryclients "order_service/internal/clients/inventory_clients"
	productclients "order_service/internal/clients/product_clients"
	"order_service/internal/entity/order"
	"order_service/internal/infrastructura/repository"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewOrderMongodb(client *mongo.Client, collection *mongo.Collection) repository.OrderRepository {
	return &OrderMongodb{client: client, collection: collection}
}

func (o *OrderMongodb) CreateOrder(ctx context.Context, req order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderID := uuid.New().String()
	var totalAmount float64
	var unavailableProducts []order.UnavailableProduct

	for _, item := range req.Items {
		product, err := inventoryclients.ProductsQuantity(ctx, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve inventory: %v", err)
		}

		if product.Quantity < item.Quantity {
			unavailableProducts = append(unavailableProducts, order.UnavailableProduct{
				ProductID:        item.ProductID,
				RequestedQuantity: item.Quantity,
				AvailableQuantity: product.Quantity,
			})
		} else {
			productPrice, err := productclients.Products(item.ProductID)
			log.Println(item.ProductID)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve product: %v", err)
			}
			if productPrice.DiscountPrice != 0{
				totalAmount += float64(productPrice.DiscountPrice) * float64(item.Quantity)
			}else{
				totalAmount += float64(productPrice.Price) * float64(item.Quantity)
			}
		}
	}

	if len(unavailableProducts) > 0 {
		return &order.CreateOrderResponse{
			Status:             "failure",
			Message:            "Some products are out of stock",
			Product: unavailableProducts,
		}, nil
	}
	orderr := order.Order{
		OrderId:      orderID,
		UserID:       req.UserID,
		Items:        req.Items,
		TotalPrice:  totalAmount,
		Status:     "true",
		CreatedAt:    time.Now(),
	}

	_, err := o.collection.InsertOne(ctx, orderr)
	if err != nil {
		return nil, fmt.Errorf("failed to insert order into MongoDB: %v", err)
	}

	return &order.CreateOrderResponse{
		Status:      "success",
		Message:     "Order created successfully",
		OrderId:     orderID,
		TotalPrice: totalAmount,
	}, nil
}


func (o *OrderMongodb) GetOrder(ctx context.Context, id string) (*order.Order, error) {
	var result order.Order
	err := o.collection.FindOne(ctx, bson.M{"order_id": id, "status" : "true"}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("order not found with ID: %s", id)
		}
		return nil, fmt.Errorf("failed to get order: %v", err)
	}
	return &result, nil
}

func (o *OrderMongodb) GetUserOrders(ctx context.Context, userID string) (*[]order.Order, error) {
	var orders []order.Order
	cursor, err := o.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order order.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, fmt.Errorf("failed to decode order: %v", err)
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func (o *OrderMongodb) GetAllOrders(ctx context.Context) (*[]order.Order, error) {
	var orders []order.Order
	cursor, err := o.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order order.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, fmt.Errorf("failed to decode order: %v", err)
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func (o *OrderMongodb) UpdateOrderStatus(ctx context.Context, orderID string, status string) (*order.Order, error) {
	filter := bson.M{"order_id": orderID}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedOrder order.Order
	err := o.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedOrder)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("order not found with ID: %s", orderID)
		}
		return nil, fmt.Errorf("failed to update order status: %v", err)
	}

	return &updatedOrder, nil
}

func (o *OrderMongodb) DeleteOrder(ctx context.Context, orderID string) error {

	result, err := o.collection.DeleteOne(ctx, bson.M{"order_id": orderID})
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("order with ID %s not found", orderID)
	}

	return nil
}
