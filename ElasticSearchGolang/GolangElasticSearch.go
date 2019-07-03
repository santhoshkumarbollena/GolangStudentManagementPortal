package main

import (

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
type ELasticSearch struct{
	_index string `json:"_index"`
	_type string `json:"_type"`
	_id string `json:"_id"`
	_version string `json:"_version"`
	_seq_no string `json:"_seq_no"`
	_primary_term string `json:"_primary_term"`
	found string `json:"found"`
	_source Student `json:"_source"`
	
}
type Student struct{
	Id string `json:"Id"`
	Name string `json:"Name"`
	Age string `json:"Age"`
	Country string `json:"Country"`
	Email string `json:"Email"`
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/GetStudent/{id}", GetStudent)
	log.Fatal(http.ListenAndServe(":1022", myRouter))

}
func GetStudent(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println()
	fmt.Println("2nd Service")
	fmt.Println()
	vars := mux.Vars(r)
	ID := vars["id"]
	fmt.Println()
	Response2ndService, _ := http.Get("http://localhost:9200/students/student/"+ID)
	defer Response2ndService.Body.Close()

	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	fmt.Println()
	fmt.Println("Printing Response")
	fmt.Println(string(ResponseBody))

	s:=string(ResponseBody)
	fmt.Println("         rep body")
	fmt.Println(s)
	var raw map[string]interface{}
	json.Unmarshal([]byte(s), &raw)
	fmt.Println("        raw")
    fmt.Println(raw)
// 	out, _ := json.Marshal(raw)
// 	fmt.Println("        out")
// 	fmt.Println(string(out))
// 	var OutputFromSecondService ELasticSearch
// 	OutputFromSecondService=ELasticSearch{_index:"students",_type:"student",_id:"2",_version:"2",_seq_no:"3",_primary_term:"2",found:"true",_source:Student{
// 		Id: "5",
// 		Name: "Bollena santhosh kumar",
// 		Age: "23",
// 		Country: "INdia" ,
// 		Email :"bollenasanthosh@gmail.com",
// 	}}
// //	json.Unmarshal([]byte(s), &OutputFromSecondService)
// 	fmt.Println("          satic obj")
// 	fmt.Println(OutputFromSecondService)
	json.NewEncoder(w).Encode(raw)
	
}

