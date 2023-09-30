package quizes

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/IraIvanishak/quiz-pet-app/config"
)

func CountPoints(answear_ids []string, id_test_i int) (int, error) {
	right_answear_ids, err := TestAnswears(id_test_i)
	if err != nil {
		return -1, err
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

func AddTestResultToSession(session_id string, test_i string, points int) {

	var dataBytes []byte
	exist := config.DB.QueryRow("SELECT results FROM usersresults WHERE sessionid = $1", session_id).Scan(&dataBytes)

	var usersResults = make(map[string]string)

	if exist == nil {
		err := json.Unmarshal(dataBytes, &usersResults)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	usersResults[test_i] = strconv.Itoa(points)

	jsonData, err := json.Marshal(usersResults)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = config.DB.Query("INSERT INTO usersresults VALUES ($1, $2) ON CONFLICT (sessionid) DO UPDATE SET results = $2", session_id, jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetAllUserResults(session_id string) (map[string]string, error) {
	var dataBytes []byte
	exist := config.DB.QueryRow("SELECT results FROM usersresults WHERE sessionid = $1", session_id).Scan(&dataBytes)
	if exist == nil {
		var usersResults = make(map[string]string)
		err := json.Unmarshal(dataBytes, &usersResults)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return usersResults, nil
	}
	return nil, exist
}
