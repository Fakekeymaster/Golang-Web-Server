package middleware

import (
    "net/http"
    "strings"
    "sync"
    "time"
)

type rateInfo struct {
    count int
    reset time.Time
}

var rateMap = make(map[string]*rateInfo)
var mu sync.Mutex

func RateLimit(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ip := strings.Split(r.RemoteAddr, ":")[0]
        now := time.Now()

        mu.Lock()
        info, exists := rateMap[ip]

        if !exists || now.After(info.reset) {
            rateMap[ip] = &rateInfo{
                count: 1,
                reset: now.Add(1 * time.Minute),
            }
            mu.Unlock()
            next(w, r)
            return
        }

        if info.count >= 5 {
            mu.Unlock()
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

        info.count++
        mu.Unlock()

        next(w, r)
    }
}
