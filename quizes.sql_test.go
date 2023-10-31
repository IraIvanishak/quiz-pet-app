package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/IraIvanishak/quiz-pet-app/storage/dbs"
	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
	"github.com/stretchr/testify/assert"
)

var queriesT *dbs.Queries
var ctxT context.Context

func TestMain(m *testing.M) {
	ctxT = context.Background()
	connStr := SetupTestDatabase()

	dbT, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	queriesT = dbs.New(dbT)

	code := m.Run()

	if dbT != nil {
		dbT.Close()
	}

	os.Exit(code)
}

func TestQuizesAll(t *testing.T) {
	res, err := queriesT.QuizesAll(ctxT)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	assert.GreaterOrEqual(t, len(res), 3)

	assert.Equal(t, res[0].Title, "History Quiz")
	assert.Equal(t, res[0].Description, "Test your knowledge of historical events.")
}

func TestQuizeByID(t *testing.T) {
	// квіз з id = 1 (Історичний квіз)
	res, err := queriesT.QuizeByID(ctxT, 1)
	assert.NoError(t, err)
	assert.Equal(t, res.Title, "History Quiz")
	assert.Equal(t, res.Description, "Test your knowledge of historical events.")
}

func TestQuestionsByTestID(t *testing.T) {
	// питання для історичного квізу з test_id = 1
	res, err := queriesT.QuestionsByTestID(ctxT, 1)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// чи є хоча б одне питання
	assert.GreaterOrEqual(t, len(res), 1)

	// конкретні дані для одного з питань
	assert.True(t, res[0].QuestionText.Valid)
	assert.Equal(t, res[0].QuestionText.String, "Who was the first President of the United States?")
	assert.NotEmpty(t, res[0].Options)
}

var uuID = uuid.New()

func TestInsertResult(t *testing.T) {
	results := make(map[int]int)
	results[0] = 2

	jsonData, _ := json.Marshal(results)
	nullRawMessage := pqtype.NullRawMessage{RawMessage: jsonData, Valid: true}

	err := queriesT.InsertResult(ctxT, dbs.InsertResultParams{Sessionid: uuID, Results: nullRawMessage})

	assert.NoError(t, err)
}

func TestResultsByID(t *testing.T) {

	res, err := queriesT.ResultsByID(ctx, uuID)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
