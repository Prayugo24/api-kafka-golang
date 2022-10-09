package config

import (
	"net/http"
)

// CORS CONFIGURATION
func CORS(next http.Handler) http.Handler{
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
		// Set Header
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// next
		next.ServeHTTP(w, r)
		return
	})
}