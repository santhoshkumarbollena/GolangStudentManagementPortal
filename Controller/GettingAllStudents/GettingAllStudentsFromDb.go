package GettingAllStudents

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Model "../../Model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
)

type Student Model.Student

func ReturnAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Methods", "OPTION")
	w.Header().Add("Content-Type", "application/json")
	//connection to cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "acheron"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()
	var Students []Student
	iter := session.Query(`SELECT id,name,age,country,email FROM student ;`).Iter()
	var Student Student
	for iter.Scan(&Student.Id, &Student.Name, &Student.Age, &Student.Country, &Student.Email) {

		fmt.Println(Student)
		Students = append(Students, Student)

	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: GetAllStudents")
	json.NewEncoder(w).Encode(Students)
}
