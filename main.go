package main

import (
	"go-cassandra/service"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func main() {

	studentService := service.NewStudentService()

	// studentService.CreateStudent()
	studentService.GetAllStudents()
	studentService.GetStudentByName("mert")

}
