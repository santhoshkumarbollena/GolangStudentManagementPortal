package GettingAllStudents

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Model "../../Model"
	_ "github.com/go-sql-driver/mysql"
)

type Student Model.Student

var Students []Student

func ReturnAllStudents(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/acheron")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("select * from student;")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var Student Student
		err = results.Scan(&Student.Id, &Student.Name, &Student.Age, &Student.Country, &Student.Email)
		if err != nil {
			panic(err.Error())
		}

		Students = append(Students, Student)
	}
	Students = append(Students, Student{Name: "Santhosh", Age: "23", Country: "India", Email: "santhu"})
	Students = append(Students, Student{Name: "Santhosh kumar bollena", Age: "23", Country: "India", Email: "santhu"})
	fmt.Println("Endpoint Hit: GetAllStudents")
	json.NewEncoder(w).Encode(Students)
}
