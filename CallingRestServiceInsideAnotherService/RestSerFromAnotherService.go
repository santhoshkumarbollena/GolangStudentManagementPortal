package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/AddStudent", AddingAStudent).Methods("POST")
	myRouter.HandleFunc("/AddAEmployee", AddingEmploye).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

type Input struct {
	RequestId                     string `json:"requestId"`
	MemberId                      string `json:"memberId"`
	MemberIdType                  string `json:"memberIdType"`
	ReferedToSpecialtyCategory    string `json:"referedToSpecialtyCategory"`
	ReferedToSpecialityCode       string `json:"referedToSpecialityCode"`
	ReferedToSpecialityAreaOfBody string `json:"referedToSpecialityAreaOfBody"`
	ProviderIds                   string `json:"providerIds"`
	SearchFilterCriteria          string `json:"searchFilterCriteria"`
	CallingApp                    string `json:"callingApp"`
	CallingAppType                string `json:"callingAppType"`
}
type Output struct {
	ResponseId     string `json:"responseId"`
	Providers      string `json:"providers"`
	NPI            string `json:"NPI"`
	Ranking        string `json:"ranking"`
	ResponseStatus string `json:"responseStatus"`
	StatusCode     string `json:"statusCode"`
	StatusMessage  string `json:"statusMessage"`
}

func AddingAStudent(w http.ResponseWriter, r *http.Request) {
	// resp, _ := http.Get("http://localhost:10000/GetAllEmployes")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)

	fmt.Println("1st Service")
	fmt.Println(Input)
	//fmt.Println(Students)
	//fmt.Println(string(reqBody))
	RequestBodyFor2ndService, _ := json.Marshal(Input)

	Response2ndService, _ := http.Post("http://localhost:10000/AddAEmployee", "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
	defer Response2ndService.Body.Close()

	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	//fmt.Println(string(ResponseBody))

	json.NewEncoder(w).Encode(string(ResponseBody))
}
func AddingEmploye(w http.ResponseWriter, r *http.Request) {
	var Output Output
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)

	fmt.Println("2nd Service")
	fmt.Println(Input)
	//Setting output values in service 2
	Output.ResponseId = "responseid"
	Output.Providers = "Providers"
	Output.NPI = "NPI"
	Output.Ranking = "Ranking"
	Output.ResponseStatus = "ResponseStatus"
	Output.StatusCode = "StatusCode"
	Output.StatusMessage = "StatusMessage"

	json.NewEncoder(w).Encode(Output)
}
