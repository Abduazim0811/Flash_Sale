package kafka

import (
	"context"
	"encoding/json"
	"log"
	"product-service/internal/entity/product"
	"product-service/internal/service"

	"github.com/twmb/franz-go/pkg/kgo"
)

type ConsumerProduct struct {
	C   *service.ProductService
}

func NewConsumer(C   *service.ProductService) ConsumerProduct{
	return ConsumerProduct{C: C}
}

func (u *ConsumerProduct) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("broker:9092"),
		kgo.ConsumeTopics("flashsale11"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	for {
		fetches := client.PollFetches(ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *ConsumerProduct) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "update":
		if err := u.Update(record.Value); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return nil
}

func (u *ConsumerProduct) Update(req []byte) error {
	var req1 product.Product

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = product.Product{
		ProductID:     req1.ProductID,
		Name:          req1.Name,
		Description:   req1.Description,
		Price:         req1.Price,
		DiscountPrice: req1.DiscountPrice,
		StockQuantity: req1.StockQuantity,
	}
	err := u.C.Updateproduct(newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
