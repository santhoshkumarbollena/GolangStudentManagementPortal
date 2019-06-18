package main

import (
	"log"
	"net/http"

	GetAllStudents "../Controller/GettingAllStudents"
	GetStudentById "../Controller/GettingStudentById"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/GetStudentById/{id}", GetStudentById.ReturnAStudent)
	myRouter.HandleFunc("/GetAllStudents", GetAllStudents.ReturnAllStudents)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

	// http.HandleFunc("/GetStudentById", GetStudentById.ReturnAStudent)
	// log.Fatal(http.ListenAndServe(":10000", nil))
	// http.HandleFunc("/GetAllStudents", GetAllStudents.ReturnAllStudents)
	// log.Fatal(http.ListenAndServe(":10000", nil))
	//GetAllStudents.HandleRequestForGettingAllStudents()
	//GetStudentById.HandleRequestForGettingStudentById()
}
