package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/IraIvanishak/questionnaire_pet/config"
	_ "github.com/lib/pq"
)

var portN = ":8080"

func main() {
	var err error
	config.DB, err = sql.Open("postgres", "host=localhost user=postgres password=coolproger dbname=questionnaire_pet sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close()
	http.HandleFunc("/", getAllTestPreview)
	http.HandleFunc("/test", getTestByID)
	fmt.Printf("Server started on port %s\n", portN)
	err = http.ListenAndServe(portN, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
