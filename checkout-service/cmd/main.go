package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/inf23056/sose-26-devops/checkout-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler)

	port := 8080
	log.Printf("Checkout Service running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
