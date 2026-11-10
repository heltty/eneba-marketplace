package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	_"github.com/jackc/pgx/v5/stdlib"
)
type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Platform string `json:"platform"`
}
func productsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.QueryContext(r.Context(),
			"SELECT id, title, platform FROM products ORDER BY id")
		if err != nil {
			log.Println("query products:", err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close();

		products := []Product{}
		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ID, &p.Title, &p.Platform); err != nil {
				log.Println("scan product:", err)
				http.Error(w, "database error", http.StatusInternalServerError)
				return
			}
			products = append(products, p)
		}
		if err := rows.Err(); err != nil {
			log.Println("rows products:", err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/products", productsHandler(db))
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}