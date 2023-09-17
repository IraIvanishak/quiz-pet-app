package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func getAllTestPreview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("SELECT id, title, description FROM tests")
	if err != nil {
		fmt.Println(err.Error())
	}

	type Test_preview struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var test_previews []Test_preview
	for rows.Next() {
		var test Test_preview
		err = rows.Scan(&test.Id, &test.Title, &test.Description)
		if err != nil {
			fmt.Println(err)
		}
		test_previews = append(test_previews, test)
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

	rows, err := db.Query("SELECT question_text, options FROM questions WHERE test_id = $1", id_test)
	if err != nil {
		fmt.Println(err)
	}
	type Option struct {
		Option_text string `json:"option_text"`
		Is_correct  bool   `json:"is_correct"`
	}

	type FrontOption struct {
		Option_text string `json:"option_text"`
		Option_id   int    `json:"option_id"`
	}

	type Question struct {
		Question_text string        `json:"question_text"`
		Options       []FrontOption `json:"options"`
	}
	type Test struct {
		Id          int        `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Questions   []Question `json:"questions"`
	}
	var questionArray []Question
	for rows.Next() {
		var optionsJson []byte
		var question Question
		err = rows.Scan(&question.Question_text, &optionsJson)
		if err != nil {
			fmt.Println(err)
		}
		var options []Option
		err = json.Unmarshal(optionsJson, &options)
		if err != nil {
			fmt.Println(err)
		}
		var frontOptions []FrontOption
		for i, val := range options {
			frontOptions = append(frontOptions, FrontOption{Option_id: i, Option_text: val.Option_text})
		}

		question.Options = frontOptions
		questionArray = append(questionArray, question)
	}

	id, _ := strconv.Atoi(id_test)
	test := Test{
		Id: id,
	}

	db.QueryRow("SELECT title, description FROM tests WHERE id = $1", id_test).Scan(&test.Title, &test.Description)
	if err != nil {
		fmt.Println(err.Error())
	}
	test.Questions = questionArray
	testJson, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	w.Write(testJson)
}

var portN = ":8080"

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=postgres password=coolproger dbname=questionnaire_pet sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	http.HandleFunc("/", getAllTestPreview)
	http.HandleFunc("/test", getTestByID)
	fmt.Printf("Server started on port %s\n", portN)
	err = http.ListenAndServe(portN, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
