package main

import (
	
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "bytes"
	"strings"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/FirstService/{requestId}/{memberId}/{memberIdType}/{referedToSpecialtyCategory}/{referedToSpecialityCode}/{referedToSpecialityAreaOfBody}/{providerIds}/{searchFilterCriteria}/{callingApp}/{callingAppType}", FirstService)
	myRouter.HandleFunc("/SecondService/{requestId}/{memberId}/{memberIdType}/{referedToSpecialtyCategory}/{referedToSpecialityCode}/{referedToSpecialityAreaOfBody}/{providerIds}/{searchFilterCriteria}/{callingApp}/{callingAppType}", SecondService)
	log.Fatal(http.ListenAndServe(":1000", myRouter))

}
func CallingServices(){
	
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

	var Input Input
vars := mux.Vars(r)
RequestId := vars["requestId"]
Input.RequestId=RequestId
MemberId := vars["memberId"]
Input.MemberId=MemberId
MemberIdType := vars["memberIdType"]	
Input.MemberIdType=MemberIdType
ReferedToSpecialtyCategory := vars["referedToSpecialtyCategory"]
Input.ReferedToSpecialtyCategory=ReferedToSpecialtyCategory
ReferedToSpecialityAreaOfBody := vars["referedToSpecialityAreaOfBody"]
Input.ReferedToSpecialityAreaOfBody=ReferedToSpecialityAreaOfBody
ReferedToSpecialityCode := vars["referedToSpecialityCode"]
ReferedToSpecialityCodeArray := strings.Split(ReferedToSpecialityCode, ",")
fmt.Println(ReferedToSpecialityCodeArray)
Input.ReferedToSpecialityCode =ReferedToSpecialityCodeArray
fmt.Println(Input.ReferedToSpecialityCode)

ProviderIds := vars["providerIds"]
ProviderIdsArray := strings.Split(ProviderIds, ",")
fmt.Println(ProviderIdsArray)
Input.ProviderIds =ProviderIdsArray


SearchFilterCriteria := vars["searchFilterCriteria"]
Input.SearchFilterCriteria=SearchFilterCriteria
CallingApp := vars["callingApp"]
Input.CallingApp=CallingApp
CallingAppType := vars["callingAppType"]
Input.CallingAppType=CallingAppType
OutputFromSecondService := Output{}
// var Input models.Input
var Output Output


fmt.Println("1st Service")
fmt.Println()
// fmt.Println(Input)
fmt.Println()
//fmt.Println(Students)
//fmt.Println(string(reqBody))

if RequestId == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: "Error no RequestId | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if MemberId == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: "Error no MemberId | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if MemberIdType == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no MemberIdType | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if ReferedToSpecialtyCategory == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no ReferedToSpecialtyCategory | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if len(ProviderIds) == 0 {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no ProviderIds | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if SearchFilterCriteria == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no SearchFilterCriteria | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if CallingApp == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no CallingApp | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}
if CallingAppType == "" {
	Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no CallingAppType | Passed"}
	json.NewEncoder(w).Encode((Output))
	return
}

fmt.Println("here2")
fmt.Println(Input)
RequestBodyFor2ndService, _ := json.Marshal(Input)
fmt.Println(string(RequestBodyFor2ndService))
Response2ndService, _ := http.Get("http://localhost:1000/SecondService/"+RequestId+"/"+MemberId+"/"+MemberIdType+"/"+ReferedToSpecialtyCategory+"/"+ReferedToSpecialityCode+"/"+ReferedToSpecialityAreaOfBody+"/"+ProviderIds+"/"+SearchFilterCriteria+"/"+CallingApp+"/"+CallingAppType)
fmt.Println(Response2ndService)
defer Response2ndService.Body.Close()
fmt.Println("here4")
ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
fmt.Println()
fmt.Println("Printing Response")
fmt.Println(string(ResponseBody))
s:=string(ResponseBody)
fmt.Println("here5")
json.Unmarshal([]byte(s), &OutputFromSecondService)

	json.NewEncoder(w).Encode(OutputFromSecondService)
}
func SecondService(w http.ResponseWriter, r *http.Request) {
	var Output Output
	
	var Input Input
	vars := mux.Vars(r)
	RequestId := vars["requestId"]
	Input.RequestId=RequestId
	MemberId := vars["memberId"]
	Input.MemberId=MemberId
	MemberIdType := vars["memberIdType"]	
	Input.MemberIdType=MemberIdType
	ReferedToSpecialtyCategory := vars["referedToSpecialtyCategory"]
	Input.ReferedToSpecialtyCategory=ReferedToSpecialtyCategory
	ReferedToSpecialityAreaOfBody := vars["referedToSpecialityAreaOfBody"]
	Input.ReferedToSpecialityAreaOfBody=ReferedToSpecialityAreaOfBody
	ReferedToSpecialityCode := vars["referedToSpecialityCode"]
	ReferedToSpecialityCodeArray := strings.Split(ReferedToSpecialityCode, ",")
	fmt.Println(ReferedToSpecialityCodeArray)
	Input.ReferedToSpecialityCode =ReferedToSpecialityCodeArray
	fmt.Println(Input.ReferedToSpecialityCode)
	
	ProviderIds := vars["providerIds"]
	ProviderIdsArray := strings.Split(ProviderIds, ",")
	fmt.Println(ProviderIdsArray)
	Input.ProviderIds =ProviderIdsArray
	
	
	SearchFilterCriteria := vars["searchFilterCriteria"]
	Input.SearchFilterCriteria=SearchFilterCriteria
	CallingApp := vars["callingApp"]
	Input.CallingApp=CallingApp
	CallingAppType := vars["callingAppType"]
	Input.CallingAppType=CallingAppType

	
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
