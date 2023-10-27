package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sqlc-dev/pqtype"

	"github.com/IraIvanishak/quiz-pet-app/config"
	"github.com/IraIvanishak/quiz-pet-app/storage/dbs"
	"github.com/google/uuid"
)

var ctx = context.Background()
var queries = dbs.New(config.DB)

func getAllTestPreview(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")
	var results map[string]string
	if err == nil {
		sessionUUID, err := uuid.Parse(sessionID.Value)
		if err != nil {
			fmt.Println(err)
		}

		dataBytes, err := queries.ResultsByID(ctx, sessionUUID)
		if err != nil {
			fmt.Println(err)
		}

		results = make(map[string]string)
		if dataBytes.Valid {
			err := json.Unmarshal([]byte(dataBytes.RawMessage), &results)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	tests, err := queries.QuizesAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	type Test_preview struct {
		Test  dbs.QuizesAllRow `json:"test"`
		Score int              `json:"score"`
	}

	var test_previews []Test_preview

	for _, test := range tests {

		var test_preview Test_preview
		test_preview.Test = test

		point, exist := results[strconv.Itoa(int(test.ID))]
		if exist {
			score, _ := strconv.Atoi(point)
			test_preview.Score = score
		} else {
			test_preview.Score = -1
		}
		test_previews = append(test_previews, test_preview)
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

	questions, err := queries.QuestionsByTestID(ctx, int32(id_test_i))
	if err != nil {
		fmt.Println(err)
	}

	testDB, err := queries.QuizeByID(ctx, int32(id_test_i))
	if err != nil {
		fmt.Println(err)
	}

	test := struct {
		Preview   dbs.QuizeByIDRow           `json:"preview"`
		Questions []dbs.QuestionsByTestIDRow `json:"questions"`
	}{
		Preview:   testDB,
		Questions: questions,
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

	// count points
	questions, err := queries.QuestionsByTestID(ctx, int32(id_test_i))
	if err != nil {
		fmt.Println(err)
	}

	var right_answear_ids []int
	for _, q := range questions {
		var options []struct {
			OptionText string `json:"option_text"`
			IsCorrect  bool   `json:"is_correct"`
		}

		err := json.Unmarshal([]byte(q.Options.RawMessage), &options)
		if err != nil {
			fmt.Println(err)
		}
		for i, o := range options {
			if o.IsCorrect == true {
				right_answear_ids = append(right_answear_ids, i)
				break
			}
		}
	}

	var points int
	for i, val := range answear_ids {
		val_i, _ := strconv.Atoi(val)
		if val_i == right_answear_ids[i] {
			points++
		}
	}

	sessionID, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println(err)

		sessionID = &http.Cookie{
			Name:     "session_id",
			Value:    uuid.New().String(),
			Path:     "/",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, sessionID)
	}

	// add test result to sesseion
	//	quizes.AddTestResultToSession(session_id.Value, id_test, points)

	sessionUUID, err := uuid.Parse(sessionID.Value)
	if err != nil {
		fmt.Println(err)
	}

	dataBytes, err := queries.ResultsByID(ctx, sessionUUID)
	if err != nil {
		fmt.Println(err)
	}

	results := make(map[string]string)
	if dataBytes.Valid {
		err := json.Unmarshal([]byte(dataBytes.RawMessage), &results)
		if err != nil {
			fmt.Println(err)
		}
	}

	results[id_test] = strconv.Itoa(points)

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}
	nullRawMessage := pqtype.NullRawMessage{RawMessage: jsonData, Valid: true}

	err = queries.InsertResult(ctx, dbs.InsertResultParams{Sessionid: sessionUUID, Results: nullRawMessage})
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write([]byte(strconv.Itoa(points)))

}

func CountPoints(answear_ids []string, id_test_i int) (int, error) {
	questions, err := queries.QuestionsByTestID(ctx, int32(id_test_i))
	if err != nil {
		fmt.Println(err)
	}

	var right_answear_ids []int
	for _, q := range questions {
		var options []struct {
			OptionText string `json:"option_text"`
			IsCorrect  bool   `json:"is_correct"`
		}

		err := json.Unmarshal([]byte(q.Options.RawMessage), &options)
		if err != nil {
			fmt.Println(err)
		}
		for i, o := range options {
			if o.IsCorrect == true {
				right_answear_ids = append(right_answear_ids, i)
				break
			}
		}
	}

	var points int
	for i, val := range answear_ids {
		val_i, _ := strconv.Atoi(val)
		if val_i == right_answear_ids[i] {
			points++
		}
	}
	return points, nil
}
