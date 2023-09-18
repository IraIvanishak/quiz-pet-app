package main

import (
	"fmt"
	"net/http"

	"github.com/IraIvanishak/quiz-pet-app/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var portN = ":8080"

func main() {
	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                               // Adjust this to specify allowed origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}, // Adjust allowed methods
		AllowedHeaders: []string{"Content-Type"},                                    // Adjust allowed headers
	})

	defer config.DB.Close()

	handler := c.Handler(http.DefaultServeMux)
	r := chi.NewRouter()

	r.Get("/", getAllTestPreview)
	r.Post("/test", getTestResult)
	r.Get("/test", getTestByID)

	http.Handle("/", r)

	fmt.Printf("Server started on port %s\n", portN)
	err := http.ListenAndServe(portN, handler)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
