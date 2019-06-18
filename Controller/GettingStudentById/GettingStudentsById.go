package GettingStudentById

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Model "../../Model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student Model.Student

func ReturnAStudent(w http.ResponseWriter, r *http.Request) {
	var Students []Student
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/acheron")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	ID := vars["id"]
	results, err := db.Query("select * from student where Id=" + ID)
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
	fmt.Println("Endpoint Hit: GetAllStudents")
	json.NewEncoder(w).Encode(Students)
}
