package main

import (
	"fmt"
	"net/http"
	"strings"
	"io"
)

//this function is run after we receive the request
//w -> where we send response
//r -> request data
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Paras\n")
	fmt.Fprintf(w, "Method: %s\n", r.Method);
	fmt.Fprintf(w, "URL is: %s\n", r.URL)
	fmt.Fprintf(w, "Headers: %v\n", r.Header)
}

func about (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about Page");
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//query is a map which give all parameters in the query after ?
	if r.Method == "GET" {
		name := r.URL.Query().Get("name")
	
		if name == "" {
			name = "guest"
		}
		fmt.Fprintf(w, "Hello %s\n", name)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		name := r.Form.Get("name")

		if name == "" {
		name = "guest"
		}
	fmt.Fprintf(w, " Post Hello %s\n", name)
	return
}
	fmt.Fprintf(w, "unsupported request type")
}

//request -> middleware -> handler
func logger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received, i am in the middleware")
		fmt.Println("Request Headers are: ", r.Header)
		fmt.Println("Ip ADDR is: ", r.RemoteAddr)
		fmt.Println("User-agent is: ", r.UserAgent())
		fmt.Println("Path is: ", r.URL.Path)

		next(w, r)
	}
}

func authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
//		token := r.Header.Get("Authorization")
		authHeader := r.Header.Get("Authorization")
		//validation
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		//validation
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization Format", http.StatusUnauthorized)
			return
		}
		
		token := parts[1]
		if token != "secret123" {
			http.Error(w, "Unautorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := "https://jsonplaceholder.typicode.com/todos/1"
	//send request
	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, "error Fetching Data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	//copy response to client
	//convert stream to usable data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Response", http.StatusInternalServerError)
	}

	w.Write(body)
}



func main() {
	//route registration '/'
	//if '/' is requested, call handler
	http.HandleFunc("/", logger(authorize(handler)))

	//if '/about' is requested, call about func
	http.HandleFunc("/about", about);

	http.HandleFunc("/hello", logger(authorize(helloHandler)))

	http.HandleFunc("/proxy", logger(proxyHandler))
	
	//Starting server
	fmt.Println("Web server is running")
	http.ListenAndServe(":8080", nil)
	

}
