package main

import (
	"fmt"
	"net/http"
	"os"

	handleapis "github.com/MarkVivian/wordGuesserHelper/handleApis"
	"github.com/joho/godotenv"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	erro := godotenv.Load(".env")

	if erro != nil {
		fmt.Printf("Error loading .env file: %v \n", erro)
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("/randomWord", handleapis.RandomWord)
	mux.HandleFunc("/searchWord", handleapis.SearchWord)

	fmt.Printf("Server listening to port %s \n ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), corsMiddleware(mux))

	if err != nil {
		fmt.Printf("Error while starting the server: %v \n", err)
	}
}
