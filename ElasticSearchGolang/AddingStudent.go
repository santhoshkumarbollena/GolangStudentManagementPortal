package main

import (

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct{
	Id string `json:"Id"`
	Name string `json:"Name"`
	Age string `json:"Age"`
	Country string `json:"Country"`
	Email string `json:"Email"`
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/AddStudent/{key}", AddStudent)
	log.Fatal(http.ListenAndServe(":1022", myRouter))

}
func AddStudent(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println()
	fmt.Println("2nd Service")
	fmt.Println()
	vars := mux.Vars(r)
	ID := vars["id"]
	fmt.Println()
  var Students []Student
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Student Student
	json.Unmarshal(reqBody, &Student)
	// update our global Articles array to include
	// our new Article
	Students = append(Students, Student)
	fmt.Println(Student)
  RequestBodyFor2ndService, _ := json.Marshal(Student)
  Response2ndService, _ := http.Put("http://localhost:9200/students/student/"+ID, "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
  
	//Response2ndService, _ := http.Put("http://localhost:9200/students/student/"+ID)
	
}
