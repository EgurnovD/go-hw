package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	FullName         string `json:"name"`
	MathScore        int    `json:"math_score"`
	InformaticsScore int    `json:"info_score"`
	EnglishScore     int    `json:"eng_score"`
}

func (st Student) String() string {
	return fmt.Sprintf("name: %s; math: %v; info: %v; eng: %v;",
		st.FullName, st.MathScore, st.InformaticsScore, st.EnglishScore)
}

var students = []Student{}

func getStudents(c *gin.Context) {
	fmt.Println("Total admitted: ", len(students))
	c.IndentedJSON(http.StatusOK, students)
}

func applyStudent(c *gin.Context) {
	var st Student

	if err := c.BindJSON(&st); err != nil {
		return
	}

	total := st.MathScore + st.InformaticsScore + st.EnglishScore
	if total >= 14 {
		students = append(students, st)
		c.IndentedJSON(http.StatusAccepted, st)
		fmt.Println("Accepted: ", st)
		return
	}
	fmt.Println("Rejected: ", st)
	c.IndentedJSON(http.StatusNotAcceptable, st)
}

func main() {
	route := gin.Default()

	route.GET("/admitted", getStudents)
	route.POST("/apply", applyStudent)

	route.Run(":8080")
}
