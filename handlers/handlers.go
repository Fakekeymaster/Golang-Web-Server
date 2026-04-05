package handlers

import (
    "fmt"
    "io"
    "net/http"
    "sync"
)

var backends = []string{
    "https://jsonplaceholder.typicode.com",
    "https://jsonplaceholder.typicode.com",
}

var counter int
var lbMutex sync.Mutex

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Paras\n")
    fmt.Fprintf(w, "Method: %s\n", r.Method)
    fmt.Fprintf(w, "URL is: %s\n", r.URL)
    fmt.Fprintf(w, "Headers: %v\n", r.Header)
}

func About(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "This is the about Page")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        name := r.URL.Query().Get("name")
        if name == "" {
            name = "guest"
        }
        fmt.Fprintf(w, "Hello %s\n", name)
        return
    }

    if r.Method == http.MethodPost {
        r.ParseForm()
        name := r.Form.Get("name")
        if name == "" {
            name = "guest"
        }
        fmt.Fprintf(w, "Post Hello %s\n", name)
        return
    }

    fmt.Fprintf(w, "unsupported request type")
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[len("/proxy"):]
    target := getNextBackend() + path
    if r.URL.RawQuery != "" {
        target += "?" + r.URL.RawQuery
    }

    req, err := http.NewRequest(r.Method, target, r.Body)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    req.Header = r.Header.Clone()

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Upstream request failed", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    for k, vals := range resp.Header {
        for _, v := range vals {
            w.Header().Add(k, v)
        }
    }

    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

func getNextBackend() string {
    lbMutex.Lock()
    defer lbMutex.Unlock()

    backend := backends[counter%len(backends)]
    counter++
    return backend
}
