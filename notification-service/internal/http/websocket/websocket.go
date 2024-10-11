package websocket

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

type FlashSaleWebSocket struct {
	clients     map[string]*websocket.Conn
	mu          sync.Mutex
	ctx         context.Context
	kafkaReader *kgo.Client
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewFlashSaleWebSocket(brokers []string) (*FlashSaleWebSocket, error) {
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeTopics("flash-sale-topic"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %v", err)
	}

	return &FlashSaleWebSocket{
		clients:     make(map[string]*websocket.Conn),
		kafkaReader: kafkaClient,
		ctx:         context.Background(),
	}, nil
}

func (ws *FlashSaleWebSocket) HandleWebSocket(c *gin.Context) {
	fmt.Println("Flash-Sale WebSocket is working")
	userID := c.Request.Header.Get("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()

	ws.mu.Lock()
	ws.clients[userID] = conn
	ws.mu.Unlock()

	defer func() {
		ws.mu.Lock()
		delete(ws.clients, userID)
		ws.mu.Unlock()
	}()

	ws.listenKafka(conn)
}

func (ws *FlashSaleWebSocket) listenKafka(conn *websocket.Conn) {
	for {
		fetches := ws.kafkaReader.PollFetches(ws.ctx)
		if fetches.IsClientClosed() {
			break
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			message := record.Value
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Error writing message to WebSocket:", err)
				return
			}
		}
	}
}

func (ws *FlashSaleWebSocket) BroadcastMessage(message string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for userID, conn := range ws.clients {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("Error sending message to user %s: %v", userID, err)
			conn.Close()
			delete(ws.clients, userID)
		}
	}
}

func (ws *FlashSaleWebSocket) Close() {
	ws.kafkaReader.Close()
}
