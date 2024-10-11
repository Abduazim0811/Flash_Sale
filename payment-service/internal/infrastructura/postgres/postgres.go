package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"payment_service/internal/clients/inventoryclients"
	"payment_service/internal/clients/orderclients"
	userclients "payment_service/internal/clients/user_clients"
	"payment_service/internal/entity/payment"
	"payment_service/internal/infrastructura/repository"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type PaymentPostgres struct {
	db *sql.DB
}

func NewPaymentPostgres(db *sql.DB) repository.PaymentRepository {
	return &PaymentPostgres{db: db}
}

func (p *PaymentPostgres) InsertPaymentPostgres(req *payment.PaymentRequest) (*payment.GetPaymentRequest, error) {
	orderRes, err := orderclients.GetOrder(req.OrderID)
	if err != nil {
		log.Println("order not found: ", err)
		return nil, fmt.Errorf("order not found: %v", err)
	}
	orders := orderRes.Order
	if req.Amount <= float32(orders.TotalPrice) {
		log.Println("not enough money: ", err)
		return nil,fmt.Errorf("not enough money: %v", err)
	}
	paymentId := uuid.New().String()
	price := req.Amount - float32(orders.TotalPrice)
	fmt.Println(price)
	err = userclients.UpdatePrice(req.UserID, price)
	if err != nil {
		log.Println("users updated error:", err)
		return nil,fmt.Errorf("users updated error: %v", err)
	}

	sql, args, err := squirrel.
		Insert("payments").
		Columns("id, order_id, user_id, amount,status").
		Values(paymentId, req.OrderID, req.UserID, req.Amount, "true").
		Suffix("Returning id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Error generating SQL for insert:", err)
		return nil,fmt.Errorf("insert error: %v", err)
	}
	var id string
	row := p.db.QueryRow(sql, args...)

	if err := row.Scan(&id); err != nil {
		log.Println("Error executing SQL insert:", err)
		return nil,fmt.Errorf("sql insert product error: %v", err)
	}
	for _, item := range orders.Items{
		err := inventoryclients.UpdateInventory(item.ProductId, item.Quantity)
		if err != nil {
			log.Println("error:")
		}
	}

	return &payment.GetPaymentRequest{PaymentID: id}, nil
}

func (p *PaymentPostgres) GetbyIDPaymentPostgres(paymentId string) (*payment.Payment, error) {
	sql, args, err := squirrel.
		Select("*").
		From("payments").
		Where(squirrel.Eq{"id": paymentId}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Error generating SQL for select:", err)
		return nil, fmt.Errorf("select error: %v", err)
	}
	var paymentRes payment.Payment

	row := p.db.QueryRow(sql, args...)
	if err := row.Scan(&paymentRes.ID, &paymentRes.OrderID, &paymentRes.UserID, &paymentRes.Amount, &paymentRes.Status, &paymentRes.CreatedAt); err != nil {
		log.Println("Error fetching payments from database:", err)
		return nil, fmt.Errorf("sql get payments error: %v", err)
	}

	return &paymentRes, nil
}
