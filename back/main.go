package main

import (
	"fmt"
	"net/http"

	"github.com/IraIvanishak/quiz-pet-app/config"
	_ "github.com/lib/pq"
)

var portN = ":8080"

func main() {
	var err error

	defer config.DB.Close()
	http.HandleFunc("/", getAllTestPreview)
	http.HandleFunc("/test", getTestByID)
	fmt.Printf("Server started on port %s\n", portN)
	err = http.ListenAndServe(portN, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
