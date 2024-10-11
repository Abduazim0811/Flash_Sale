package router

import (
	flashsaleclients "api-gateway/internal/clients/flashSale_clients"
	inventoryclients "api-gateway/internal/clients/inventory_clients"
	orderclients "api-gateway/internal/clients/order_clients"
	paymentclients "api-gateway/internal/clients/payment_clients"
	productclients "api-gateway/internal/clients/product_clients"
	userclients "api-gateway/internal/clients/user_clients"
	"api-gateway/internal/config"
	flashsalehandler "api-gateway/internal/https/api/handlers/flashSale_handler"
	inventoryhandler "api-gateway/internal/https/api/handlers/inventory_handler"
	orderhandler "api-gateway/internal/https/api/handlers/order_handler"
	paymentshandler "api-gateway/internal/https/api/handlers/payments_handler"
	productshandler "api-gateway/internal/https/api/handlers/products_handler"
	usershandler "api-gateway/internal/https/api/handlers/users_handler"
	"api-gateway/internal/pkg/jwt"
	middleware "api-gateway/internal/rate-limiting"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *http.Server {
	userclient := userclients.DialUsersGrpc()
	userhandler := &usershandler.UsersHandler{ClientUser: userclient}

	productclient := productclients.DilaProductGrpc()
	producthandler := &productshandler.ProductsHandler{ClientProduct: productclient}

	inventoryclient := inventoryclients.DialInventoryGrpc()
	inventoryhandler := &inventoryhandler.InventoryHandler{ClientInventory: inventoryclient}

	orderclient := orderclients.DialOrderGrpc()
	orderhandler := &orderhandler.OrdersHandler{ClientOrder: orderclient}

	paymentclient := paymentclients.DialPaymentGrpc()
	paymenthandler := &paymentshandler.PaymentsHandler{ClientPayment: paymentclient}

	flashsalesclient := flashsaleclients.DialFlashSaleGrpc()
	flashsaleshandler := &flashsalehandler.FlashSalesHandler{ClientFlashSale: flashsalesclient}

	userRateLimiter := middleware.NewRateLimiter(5, time.Minute)
	productRateLimiter := middleware.NewRateLimiter(4, time.Minute)
	orderRateLimiter := middleware.NewRateLimiter(4, time.Minute)
	inventoryRateLimiter := middleware.NewRateLimiter(4, time.Minute)
	paymentRateLimiter := middleware.NewRateLimiter(4, time.Minute)
	flashsaleRateLimiter := middleware.NewRateLimiter(3, time.Minute)

	router := gin.Default()
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userRoutes := router.Group("/")
	userRoutes.Use(userRateLimiter.Limit())
	{
		userRoutes.POST("register", userhandler.Register)
		userRoutes.POST("verify-code", userhandler.VerifyCode)
		userRoutes.POST("login", userhandler.Login)
		userRoutes.GET("users/:id", jwt.Protected(), userhandler.GetbyIdUsers)
		userRoutes.GET("users", jwt.Protected(),userhandler.GetAllUsers)
		userRoutes.PUT("users/:id", jwt.Protected(),userhandler.UpdateUsers)
		userRoutes.PUT("users/:id/password", jwt.Protected(),userhandler.UpdateUsersPassword)
		userRoutes.DELETE("users/:id", jwt.Protected(),userhandler.DeleteUsers)
	}

	productRoutes := router.Group("/")
	productRoutes.Use(productRateLimiter.Limit())
	{
		productRoutes.POST("product", jwt.Protected(), producthandler.Createproduct)
		productRoutes.GET("product/:id", jwt.Protected(), producthandler.GetbyIdProduct)
		productRoutes.GET("products", jwt.Protected(), producthandler.GetAllProduct)
		productRoutes.PUT("product/:id", jwt.Protected(), producthandler.Updateproduct)
		productRoutes.DELETE("product/:id", jwt.Protected(), producthandler.Deleteproduct)
	}

	inventoryRoutes := router.Group("/")
	inventoryRoutes.Use(inventoryRateLimiter.Limit())
	{
		inventoryRoutes.POST("inventory", jwt.Protected(), inventoryhandler.Createinventory)
		inventoryRoutes.GET("inventory/:id", jwt.Protected(), inventoryhandler.GetbyIdinventory)
		inventoryRoutes.GET("inventory", jwt.Protected(), inventoryhandler.GetAllinventory)
		inventoryRoutes.PUT("inventory", jwt.Protected(), inventoryhandler.Updateinventory)
	}

	orderRoutes := router.Group("/")
	orderRoutes.Use(orderRateLimiter.Limit())
	{
		orderRoutes.POST("order/create", jwt.Protected(), orderhandler.Createorder)
		orderRoutes.GET("order/user/:id", jwt.Protected(), orderhandler.GetUserOrders)
		orderRoutes.GET("order/:id", jwt.Protected(), orderhandler.GetbyIdOrder)
		orderRoutes.GET("orders", jwt.Protected(), orderhandler.GetAllOrders)
		orderRoutes.PUT("order/status", jwt.Protected(), orderhandler.UpdateOrderstatus)
		orderRoutes.DELETE("order/:id", jwt.Protected(), orderhandler.Deleteorder)
	}

	paymentRoutes := router.Group("/")
	paymentRoutes.Use(paymentRateLimiter.Limit())
	{
		paymentRoutes.POST("payment/process", jwt.Protected(),paymenthandler.ProcessPayments)
		paymentRoutes.GET("payment/:id", jwt.Protected(),paymenthandler.Getpayment)
	}

	flashsaleRoutes := router.Group("/")
	flashsaleRoutes.Use(flashsaleRateLimiter.Limit())
	{
		flashsaleRoutes.POST("flashsale", jwt.Protected(), flashsaleshandler.CreateFlashsale)
		flashsaleRoutes.GET("flashsale", jwt.Protected(), flashsaleshandler.GetAllFlashSales)
		flashsaleRoutes.GET("flashsale/active", jwt.Protected(), flashsaleshandler.GetActiveFlashSales)
		flashsaleRoutes.PUT("flashsale", jwt.Protected(), flashsaleshandler.UpdateFlashsale)
		flashsaleRoutes.DELETE("/flashsale", jwt.Protected(), flashsaleshandler.DeleteFlashsale)
		flashsaleRoutes.POST("flashsale/purchase", jwt.Protected(), flashsaleshandler.Purchaseproduct)
	}
	c := config.Configuration()
	server := &http.Server{
		Addr:    c.ApiGateway.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServeTLS("./internal/tls/items.pem", "./internal/tls/items-key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run HTTPS server: %v", err)
		}
	}()

	GracefulShutdown(server, log.Default())

	return server

}

func GracefulShutdown(srv *http.Server, logger *log.Logger) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	<-shutdownCh
	logger.Println("Shutdown signal received, initiating graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Printf("Server shutdown encountered an error: %v", err)
	} else {
		logger.Println("Server gracefully stopped")
	}

	select {
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			logger.Println("Shutdown deadline exceeded, forcing server to stop")
		}
	default:
		logger.Println("Shutdown completed within the timeout period")
	}
}
