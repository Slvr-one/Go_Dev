package test

import (
	"fmt"
	"net/http"
)

func main() {
	// converting our handler function to handler
	// type to make use of our middleware
	myHandler := http.Handlerfunc(handler)
	http.Handle("/", middleware(myHandler)) // 👈
	http.HandleAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Prinln("Executing the handler")
	w.Write([]byte("OK"))
}

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		w.Write([]byte("Hijacking Request "))
		originalHandler.Serve(w, r)
		fmt.Println("Running after handler")
	})
}
