package main

import (
    "fmt"
    "io/ioutil"
	"os"
	"encoding/json"
	
	"log"
	"net/http"

	"github.com/gorilla/mux"
    
)

func main() {
   

  myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/GetStudentById", ReturnAStudent)

  log.Fatal(http.ListenAndServe(":10000", myRouter))
}

type Student struct{
	 Name string `json:"Name"`
	 Id string `json:"Id"`
}
func ReturnAStudent(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./JsonData.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


  b, err := ioutil.ReadAll(file)
  fmt.Print(string(b))
  s:=string(b)
  OutputFromSecondService := Student{}
//	w.Header().Set("Access-Control-Allow-Origin", "*")

//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// var Students []Student
	// cluster := gocql.NewCluster("127.0.0.1")
	// cluster.Keyspace = "acheron"
	// cluster.Consistency = gocql.Quorum
	// session, _ := cluster.CreateSession()
	// defer session.Close()

	// vars := mux.Vars(r)
	// ID := vars["id"]
	// var Student Student
	// if err := session.Query(`SELECT id,name,age,country,email FROM student WHERE id = ? LIMIT 1`,
	// 	ID).Consistency(gocql.One).Scan(&Student.Id, &Student.Name, &Student.Age, &Student.Country, &Student.Email); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(Student)
	// Students = append(Students, Student)
	// var student Student
	// json.Unmarshal([]byte(s), &student)
	// fmt.Println("Endpoint Hit: GetStudentsById")
	// fmt.Println(student)
	
	
	json.Unmarshal([]byte(s), &OutputFromSecondService)
	fmt.Println(OutputFromSecondService)
	json.NewEncoder(w).Encode(OutputFromSecondService)
}
