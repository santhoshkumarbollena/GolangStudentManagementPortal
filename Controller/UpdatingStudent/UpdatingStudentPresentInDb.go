package UpdatingStudent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	Model "../../Model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
)

type Student Model.Student

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, OPTIONS, POST, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	var Students []Student
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Student Student
	json.Unmarshal(reqBody, &Student)
	// update our global Articles array to include
	// our new Article
	Students = append(Students, Student)
	fmt.Println(Student)
	json.NewEncoder(w).Encode(Student)

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "acheron"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	//insert a tweet
	if err := session.Query(`INSERT INTO  student (id,name,age,country,email) VALUES (?, ?, ?,?,?)`,
		Student.Id, Student.Name, Student.Age, Student.Country, Student.Email).Exec(); err != nil {
		log.Fatal(err)
	}
}
