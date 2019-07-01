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

	myRouter.HandleFunc("/FirstService", FirstService).Methods("POST")
	myRouter.HandleFunc("/SecondService", SecondService).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

type Input struct {
	RequestId                     string   `json:"requestId"`
	MemberId                      string   `json:"memberId"`
	MemberIdType                  string   `json:"memberIdType"`
	ReferedToSpecialtyCategory    string   `json:"referedToSpecialtyCategory"`
	ReferedToSpecialityCode       []string `json:"referedToSpecialityCode"`
	ReferedToSpecialityAreaOfBody string   `json:"referedToSpecialityAreaOfBody"`
	ProviderIds                   []string `json:"providerIds"`
	SearchFilterCriteria          string   `json:"searchFilterCriteria"`
	CallingApp                    string   `json:"callingApp"`
	CallingAppType                string   `json:"callingAppType"`
}
type Providers struct {
	NPI            string      
	Ranking        string   
}
type Object struct {
	StatusCode     string     
	StatusMessage  string 
}
type Output struct {
	ResponseId     string      `json:"responseId"`
	Providers      []Providers `json:"providers"`
	  
	ResponseStatus Object      `json:"responseStatus"`
    
}

func FirstService(w http.ResponseWriter, r *http.Request) {
	// resp, _ := http.Get("http://localhost:10000/GetAllEmployes")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	reqBody, _ := ioutil.ReadAll(r.Body)
	OutputFromSecondService := Output{}
	var Input Input
	var Output Output
	json.Unmarshal(reqBody, &Input)

	fmt.Println("1st Service")
	fmt.Println()
	fmt.Println(Input)
	fmt.Println()
	//fmt.Println(Students)
	//fmt.Println(string(reqBody))

	if(Input.RequestId==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:"Error no RequestId | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(Input.MemberId==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:"Error no MemberId | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(Input.MemberIdType==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:" Error no MemberIdType | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(Input.ReferedToSpecialtyCategory==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:" Error no ReferedToSpecialtyCategory | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(len(Input.ProviderIds)==0){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:" Error no ProviderIds | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(Input.SearchFilterCriteria==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:" Error no SearchFilterCriteria | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(Input.CallingApp==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:" Error no CallingApp | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if(Input.CallingAppType==""){
		Output.ResponseStatus = Object{StatusCode:"901",StatusMessage:" Error no CallingAppType | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	RequestBodyFor2ndService, _ := json.Marshal(Input)

	Response2ndService, _ := http.Post("http://localhost:10000/SecondService", "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
	defer Response2ndService.Body.Close()

	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	fmt.Println()
	fmt.Println("Printing Response")
	fmt.Println(string(ResponseBody))
	s:=string(ResponseBody)
	
	json.Unmarshal([]byte(s), &OutputFromSecondService)
	json.NewEncoder(w).Encode(OutputFromSecondService)
}
func SecondService(w http.ResponseWriter, r *http.Request) {
	var Output Output
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)
	
	fmt.Println()
	fmt.Println("2nd Service")
	fmt.Println()
	fmt.Println(Input)
	fmt.Println()
	//Setting output values in service 2
	Output.ResponseId = "1231"
	Provider := Providers{"key1","1"}
	//Object := Object{"Response Code","Response Status"}
	var ProvidersList []Providers
	ProvidersList = append(ProvidersList, Provider)
	Provider2 := Providers{"key2","2"}
	ProvidersList = append(ProvidersList, Provider2)
	Output.Providers = ProvidersList
	
	Output.ResponseStatus = Object{StatusCode:"200",StatusMessage:"sucess"}


	json.NewEncoder(w).Encode(Output)
}
