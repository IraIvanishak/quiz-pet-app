package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IraIvanishak/quiz-pet-app/models/quizes"
	"github.com/google/uuid"
)

func getAllTestPreview(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("session_id")
	var results map[string]string
	if err == nil {
		results, err = quizes.GetAllUserResults(session_id.Value)
		if err != nil {
			fmt.Println(err)
		}
	}

	test_previews, err := quizes.AllTestPreview(results)
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

	var answear_ids []string
	err := json.NewDecoder(r.Body).Decode(&answear_ids)
	if err != nil {
		fmt.Println(err)
	}

	points, err := quizes.CountPoints(answear_ids, id_test_i)
	if err != nil {
		fmt.Println(err)
	}

	session_id, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println(err)

		session_id = &http.Cookie{
			Name:     "session_id",
			Value:    uuid.New().String(),
			Path:     "/",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, session_id)
	}
	quizes.AddTestResultToSession(session_id.Value, id_test, points)

	w.Write([]byte(strconv.Itoa(points)))

}
