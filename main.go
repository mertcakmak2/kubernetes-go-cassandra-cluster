package main

import (
	"go-cassandra/model"
	"go-cassandra/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func main() {

	studentService := service.NewStudentService()

	router := gin.Default()

	router.GET("/api/students", func(ctx *gin.Context) {
		ctx.JSON(200, studentService.GetAllStudents())
	})

	router.POST("/api/students", func(ctx *gin.Context) {

		var student model.Student
		if err := ctx.ShouldBindJSON(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, studentService.CreateStudent(student))
	})

	router.Run(":8080")

}
