package connections

import (
	"database/sql"
	"fmt"
	"log"
	"product-service/internal/config"
	"product-service/internal/infrastructura/kafka"

	_ "github.com/lib/pq"
)

func Database() *sql.DB {
	c := config.Configuration()

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBname))
	if err != nil {
		log.Println(err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func KafkaConnection() (*kafka.KafkaProducer, error){
	c := config.Configuration()
	brokers := []string{c.Kafka.Port}
	kafkaProducer, err := kafka.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
		return nil, fmt.Errorf("failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()
	return kafkaProducer, nil
}
