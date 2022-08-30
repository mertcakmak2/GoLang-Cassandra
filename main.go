package main

import (
	"fmt"
	"go-cassandra/model"
	"math/rand"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func main() {

	var err error
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "restfulapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra well initialized")

	createStudent()
	getAllStudents()
	// deleteStudentById(23)
	// UpdateStudent(73)

}

func createStudent() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(200-1) + 1

	st := model.Student{id, "mert", "cakmak", 23}

	if err := Session.Query("INSERT INTO students(id, firstname, lastname, age) VALUES(?, ?, ?, ?)",
		st.ID, st.Firstname, st.Lastname, st.Age).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}
}

func getAllStudents() {
	var students []model.Student
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM students").Iter()
	for iter.MapScan(m) {
		students = append(students, model.Student{
			ID:        m["id"].(int),
			Firstname: m["firstname"].(string),
			Lastname:  m["lastname"].(string),
			Age:       m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	fmt.Println(students)
}

func deleteStudentById(id int) {

	if err := Session.Query("DELETE FROM students WHERE id = ?", id).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Println("delete successfully: ", id)
}

func UpdateStudent(id int) {
	updateStudent := model.Student{id, "mert", "cakmak", 26}
	if err := Session.Query("UPDATE students SET firstname = ?, lastname = ?, age = ? WHERE id = ?",
		updateStudent.Firstname, updateStudent.Lastname, updateStudent.Age, id).Exec(); err != nil {
		fmt.Println("Error while updating")
		fmt.Println(err)
	}
	fmt.Println("updated successfully: ", updateStudent)

}
