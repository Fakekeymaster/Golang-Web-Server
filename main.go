package main

import (
    "fmt"
    "net/http"

    "go-api-gateway/handlers"
    "go-api-gateway/middleware"
)

func main() {
    http.HandleFunc("/", middleware.Logger(middleware.Authorize(handlers.Handler)))
    http.HandleFunc("/about", handlers.About)
    http.HandleFunc("/hello", middleware.LatencyLogger(middleware.RateLimit(middleware.Logger(middleware.Authorize(handlers.HelloHandler)))))
    http.HandleFunc("/proxy/", middleware.Logger(middleware.Authorize(handlers.ProxyHandler)))

    fmt.Println("Web server is running")
    http.ListenAndServe(":8080", nil)
}
