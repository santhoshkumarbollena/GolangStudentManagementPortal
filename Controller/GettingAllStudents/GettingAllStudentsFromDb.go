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

var Students []Student

func ReturnAllStudents(w http.ResponseWriter, r *http.Request) {
	//connection to cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "acheron"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	iter := session.Query(`SELECT id,name,age,country,email FROM student ;`).Iter()
	var Student Student
	for iter.Scan(&Student.Id, &Student.Name, &Student.Age, &Student.Country, &Student.Email) {
		// fmt.Println("Student:", id, name, age, country, email)
		fmt.Println(Student)
		Students = append(Students, Student)

	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// for results.Next() {
	// 	var Student Student
	// 	err = results.Scan(&Student.Id, &Student.Name, &Student.Age, &Student.Country, &Student.Email)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	Students = append(Students, Student)
	// }
	// Students = append(Students, Student{Name: "Santhosh", Age: "23", Country: "India", Email: "santhu"})
	// Students = append(Students, Student{Name: "Santhosh kumar bollena", Age: "23", Country: "India", Email: "santhu"})
	fmt.Println("Endpoint Hit: GetAllStudents")
	json.NewEncoder(w).Encode(Students)
}
