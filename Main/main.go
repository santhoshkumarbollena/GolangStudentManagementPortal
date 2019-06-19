package main

import (
	"log"
	"net/http"

	DeleteStudent "../Controller/DeletingStudent"
	GetAllStudents "../Controller/GettingAllStudents"
	GetStudentById "../Controller/GettingStudentById"
	AddStudent "../Controller/InsertingStudent"
	UpdateStudent "../Controller/UpdatingStudent"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/GetStudentById/{id}", GetStudentById.ReturnAStudent)
	myRouter.HandleFunc("/GetAllStudents", GetAllStudents.ReturnAllStudents)
	myRouter.HandleFunc("/AddStudent", AddStudent.CreateNewStudent).Methods("POST")
	myRouter.HandleFunc("/UpdateStudent", UpdateStudent.UpdateStudent).Methods("PUT")
	myRouter.HandleFunc("/DeleteStudent/{id}", DeleteStudent.DeleteAStudent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

	// http.HandleFunc("/GetStudentById", GetStudentById.ReturnAStudent)
	// log.Fatal(http.ListenAndServe(":10000", nil))
	// http.HandleFunc("/GetAllStudents", GetAllStudents.ReturnAllStudents)
	// log.Fatal(http.ListenAndServe(":10000", nil))
	//GetAllStudents.HandleRequestForGettingAllStudents()
	//GetStudentById.HandleRequestForGettingStudentById()
}
