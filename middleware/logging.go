package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received, i am in the middleware")
		fmt.Println("Request Headers are:", r.Header)
		fmt.Println("Ip ADDR is:", r.RemoteAddr)
		fmt.Println("User-agent is:", r.UserAgent())
		fmt.Println("Path is:", r.URL.Path)

		next(w, r)
	}
}

func LatencyLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		fmt.Println("Latency:", duration)
	}
}
