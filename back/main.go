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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500"}, // Specify the allowed origin here
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	defer config.DB.Close()

	r := chi.NewRouter()
	r.Use(c.Handler)

	r.Get("/", getAllTestPreview)
	r.Post("/test", getTestResult)
	r.Get("/test", getTestByID)
	r.Get("/test-res", getTestByID)

	fmt.Printf("Server started on port %s\n", portN)
	err := http.ListenAndServe(portN, r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
