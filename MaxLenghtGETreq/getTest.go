package main

import (
	
	"encoding/json"

	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/FirstService/{id}", FirstService)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

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
	
	json.NewEncoder(w).Encode("OutputFromSecondService")
}