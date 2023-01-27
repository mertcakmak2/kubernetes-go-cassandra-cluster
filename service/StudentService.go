package service

import (
	"fmt"
	"go-cassandra/model"
	"math/rand"
	"time"

	"github.com/gocql/gocql"
	"github.com/pborman/uuid"
)

type StudentService struct {
	Session *gocql.Session
}

func NewStudentService() StudentService {
	var err error
	// Kubernetes Cluster
	cluster := gocql.NewCluster("34.118.77.21:32682")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "DQpR9uDXE2"}
	cluster.Keyspace = "csdb"
	Session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra well initialized")
	return StudentService{Session: Session}
}

func (service StudentService) CreateStudent(st model.Student) model.Student {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(200-1) + 1

	if err := service.Session.Query("INSERT INTO students(id, firstname, lastname, age, lastdate) VALUES(?, ?, ?, ?, ?)",
		id, st.Firstname, st.Lastname, st.Age, time.Now()).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}

	return st
}

func (service StudentService) GetAllStudents() []model.Student {

	generateUniqueTimeBasedUUID()
	var students []model.Student
	m := map[string]interface{}{}

	iter := service.Session.Query("SELECT * FROM students").Iter()
	for iter.MapScan(m) {
		students = append(students, model.Student{
			ID:        m["id"].(int),
			Firstname: m["firstname"].(string),
			Lastname:  m["lastname"].(string),
			Age:       m["age"].(int),
			Lastdate:  m["lastdate"].(time.Time),
		})
		m = map[string]interface{}{}
	}

	fmt.Println(students)
	return students
}

func (service StudentService) GetStudentByName(name string) []model.Student {
	var students []model.Student
	m := map[string]interface{}{}

	// nw, _ := time.Now().MarshalText()
	// fmt.Println(string(nw))
	// p := time.Now().Add(-2 * time.Hour)
	// fmt.Println(p)

	query := fmt.Sprintf("SELECT * FROM students where lastdate>='%s' and firstname=? ALLOW FILTERING", "2022-08-31T11:07:06.365Z")
	iter := service.Session.Query(query, name).Iter()
	for iter.MapScan(m) {
		students = append(students, model.Student{
			ID:        m["id"].(int),
			Firstname: m["firstname"].(string),
			Lastname:  m["lastname"].(string),
			Age:       m["age"].(int),
			Lastdate:  m["lastdate"].(time.Time),
		})
		m = map[string]interface{}{}
	}
	fmt.Println("Get student by name: ", students)
	// v, _ := students[0].Lastdate.UTC().MarshalText()
	// fmt.Println(string(v))
	return students
}

func (service StudentService) deleteStudentById(id int) int {

	if err := service.Session.Query("DELETE FROM students WHERE id = ?", id).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Println("delete successfully: ", id)
	return id
}

func (service StudentService) UpdateStudent(id int) model.Student {
	updateStudent := model.Student{id, "mert", "cakmak", 26, time.Now()}
	if err := service.Session.Query("UPDATE students SET firstname = ?, lastname = ?, age = ? WHERE id = ?",
		updateStudent.Firstname, updateStudent.Lastname, updateStudent.Age, id).Exec(); err != nil {
		fmt.Println("Error while updating")
		fmt.Println(err)
	}
	fmt.Println("updated successfully: ", updateStudent)
	return updateStudent
}

func generateUniqueTimeBasedUUID() (id string) {
	id = uuid.NewUUID().String()
	return
}
