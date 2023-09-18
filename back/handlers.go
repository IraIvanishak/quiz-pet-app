package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IraIvanishak/quiz-pet-app/models/quizes"
)

func getAllTestPreview(w http.ResponseWriter, r *http.Request) {
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

func getTestResult(w http.ResponseWriter, r *http.Request) {
	id_test := r.URL.Query().Get("id")
	id_test_i, _ := strconv.Atoi(id_test)

	// checking the result
	var points int

	var answear_ids []string
	err := json.NewDecoder(r.Body).Decode(&answear_ids)
	if err != nil {
		fmt.Println(err)
	}
	right_answear_ids, err := quizes.TestAnswears(id_test_i)
	if err != nil {
		fmt.Println(err)
	}

	for i, val := range answear_ids {
		val_i, _ := strconv.Atoi(val)
		if val_i == right_answear_ids[i] {
			points++
		}
	}

	fmt.Println(points)
}
