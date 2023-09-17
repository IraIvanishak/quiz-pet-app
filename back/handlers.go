package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IraIvanishak/quiz-pet-app/models/quizes"
)

func getAllTestPreview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	test_previews, err := quizes.AllTestPreview()
	if err != nil {
		fmt.Println(err)
	}
	test_json, err := json.Marshal(test_previews)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(test_json)
}

func getTestByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id_test := r.URL.Query().Get("id")
	id_test_i, _ := strconv.Atoi(id_test)

	test, err := quizes.TestByID(id_test_i)
	if err != nil {
		fmt.Println(err)
	}
	testJson, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(testJson)
}
