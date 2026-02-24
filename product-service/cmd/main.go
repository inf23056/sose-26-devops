package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/inf23056/sose-26-devops/product-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", handler.ProductListHandler)
	mux.HandleFunc("/products/{id}", handler.ProductDetailHandler)

	port := 8080
	log.Printf("Product Service running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
