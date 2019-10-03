package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ctram/student_scores/student_score"
  "net/http"
)

func SetupStudentScoresRouter(r *gin.Engine, studentScoresMap map[string]student_score.StudentScore) *gin.Engine {
  // Ping test
	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "asd"})
	})

	return r
}
