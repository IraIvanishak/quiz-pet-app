package main

import (
	"fmt"
	"net/http"

	"github.com/IraIvanishak/quiz-pet-app/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var portN = ":8080"

const maxRetries = 10

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
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

	fmt.Printf("Server is started on port %s\n", portN)
	err := http.ListenAndServe(portN, r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
