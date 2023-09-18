package quizes

import (
	"encoding/json"
	"fmt"

	"github.com/IraIvanishak/quiz-pet-app/config"
)

type Test_preview struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Test struct {
	Preview   Test_preview `json:"preview"`
	Questions []Question   `json:"questions"`
}

type Question struct {
	Question_text string   `json:"question_text"`
	Options       []Option `json:"options"`
}

type Option struct {
	Option_id   int    `json:"option_id"`
	Option_text string `json:"option_text"`
	Is_correct  bool   `json:"is_correct"`
}

func AllTestPreview() ([]Test_preview, error) {
	rows, err := config.DB.Query("SELECT id, title, description FROM tests")
	if err != nil {
		fmt.Println(err)
		return nil, err
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

	return test_previews, nil
}

func TestByID(id_test int) (Test, error) {

	rows, err := config.DB.Query("SELECT question_text, options FROM questions WHERE test_id = $1", id_test)
	if err != nil {
		fmt.Println(err)
		return Test{}, err
	}

	var questionArray []Question
	for rows.Next() {
		var optionsJson []byte
		var question Question
		err = rows.Scan(&question.Question_text, &optionsJson)
		if err != nil {
			fmt.Println(err)
			return Test{}, err
		}
		var options []Option
		err = json.Unmarshal(optionsJson, &options)
		if err != nil {
			fmt.Println(err)
			return Test{}, err
		}
		var frontOptions []Option

		for i, val := range options {
			frontOptions = append(frontOptions, Option{
				Option_id:   i,
				Option_text: val.Option_text,
				Is_correct:  val.Is_correct,
			})
		}

		question.Options = frontOptions
		questionArray = append(questionArray, question)
	}

	test_preview := Test_preview{
		Id: id_test,
	}

	config.DB.QueryRow("SELECT title, description FROM tests WHERE id = $1", id_test).Scan(&test_preview.Title, &test_preview.Description)
	if err != nil {
		fmt.Println(err)
		return Test{}, err
	}
	test := Test{
		Preview: test_preview,
	}
	test.Questions = questionArray
	return test, nil
}

func TestAnswears(id_test int) ([]int, error) {
	test, err := TestByID(id_test)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var answears []int
	for _, q := range test.Questions {
		for i, o := range q.Options {
			if o.Is_correct == true {
				answears = append(answears, i)
			}
		}
	}

	return answears, nil
}
