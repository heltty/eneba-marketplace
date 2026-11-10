package main

import (
	"encoding/json"
	"log"
	"net/http"
)
type ProductItem struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Platform string `json:"platform"`
	PriceCents int `json:"priceCents"`
}
func productHandler(w http.ResponseWriter, r *http.Request){
	items := []ProductItem{{1, "Elden Ring", "PC", 6000}, {2, "Hades", "PlayStation 4", 2500}, {3, "Stardew Valley", "Nintendo Switch", 1500}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/products", productHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}