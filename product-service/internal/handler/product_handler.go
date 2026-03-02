package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	helpers "github.com/inf23056/sose-26-devops/product-service/pkg/helpers"
)

var products = []helpers.Product{
	{ID: 1, Name: "NPC Energy", Price: 10.00},
	{ID: 2, Name: "Cosmic Calm", Price: 16.99},
	{ID: 3, Name: "Painkillaz", Price: 8.99},
	{ID: 4, Name: "Streetsugar", Price: 39.99},
	{ID: 5, Name: "Aura-Oil", Price: 3.99},
	{ID: 6, Name: "Focus+", Price: 8.99},
	{ID: 7, Name: "Chill Puff", Price: 13.99},
	{ID: 8, Name: "Partypillz", Price: 6.99},
	{ID: 9, Name: "Magic Dust", Price: 30.99},
	{ID: 10, Name: "Moonrocks", Price: 35.99},
}

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}
	w.Write(response)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, `{"error":"Product ID has wrong format"}`, http.StatusBadRequest)
		return
	}
	product := helpers.FindProductByID(products, id)
	if product == nil {
		http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
		return
	}
	resp, err := json.Marshal(product)
	if err != nil {
		http.Error(w, `{"error":"Internal Server Error"}`, http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
