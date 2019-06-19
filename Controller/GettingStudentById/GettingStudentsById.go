package GettingStudentById

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Model "../../Model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type Student Model.Student

func ReturnAStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var Students []Student
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "acheron"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	vars := mux.Vars(r)
	ID := vars["id"]
	var Student Student
	if err := session.Query(`SELECT id,name,age,country,email FROM student WHERE id = ? LIMIT 1`,
		ID).Consistency(gocql.One).Scan(&Student.Id, &Student.Name, &Student.Age, &Student.Country, &Student.Email); err != nil {
		log.Fatal(err)
	}
	fmt.Println(Student)
	Students = append(Students, Student)

	fmt.Println("Endpoint Hit: GetStudentsById")
	json.NewEncoder(w).Encode(Students)
}
