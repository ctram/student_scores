package student_scores_caller

import (
  "net/http"
  "fmt"
  "bufio"
  "log"
  "encoding/json"
  "errors"
  "bytes"
  "github.com/ctram/student_scores/student_score"
)

func getScores(studentScoresMap map[string]student_score.StudentScore) (error) {
  url := "http://live-test-scores.herokuapp.com/scores"
  resp, err := http.Get(url)

  if err != nil {
    error := fmt.Sprintf("Error fetching scores: %v", err)
    log.Println(string(error))
    return errors.New(error)
  }

  defer resp.Body.Close()

  reader := bufio.NewReader(resp.Body)


  var studentScore student_score.StudentScore

  for {
    myByte, err := reader.ReadBytes('\n')

    if err != nil {
      error := fmt.Sprintf("Error reading server sent event: %v", err)
      log.Println(string(error))
    } else if bytes.Contains(myByte, []byte("{")) {
      myJson := bytes.Split(myByte, []byte("data: "))
      fmt.Println(string(myJson[1]))
      json.Unmarshal(myJson[1], &studentScore)
      studentScoresMap[studentScore.StudentId] = studentScore
      fmt.Println(studentScoresMap)
    }
  }
}

func SetUpStudentScoresStore() (map[string]student_score.StudentScore) {
	studentScoresMap := map[string]student_score.StudentScore{}
	getScores(studentScoresMap)
	return studentScoresMap
}
