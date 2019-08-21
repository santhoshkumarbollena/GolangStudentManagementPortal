package main

import (
	
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	// "os"
	"io"
	"net/http"
	"bytes"
	"github.com/gorilla/mux"
	"time"
	"gopkg.in/yaml.v2"
)
var ApplicationLogs[] string

var data string
type T struct {
	PortNumber string

}
func readData() {
	b, err := ioutil.ReadFile("Properties.yaml") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	data = string(b) // convert content to a 'string'

}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/FirstService", FirstService).Methods("POST")
	myRouter.HandleFunc("/SecondService", SecondService).Methods("POST")
	readData()

	m := make(map[string]string)

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//Assigning Properties data into a map
	fmt.Printf("--- m:\n%v\n\n", m)
	fmt.Println(m["portNumber"])

	log.Fatal(http.ListenAndServe(m["portNumber"], myRouter))

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
type Logger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	flag   int        // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
}
var ApplicationLogs12[] string

func FirstService(w http.ResponseWriter, r *http.Request) {
	// resp, _ := http.Get("http://localhost:10000/GetAllEmployes")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	DateTimeInMilliseconds := time.Now()
	reqBody, _ := ioutil.ReadAll(r.Body)
	OutputFromSecondService := Output{}
	var Input Input
	var Output Output
	json.Unmarshal(reqBody, &Input)
	
	
	
	// f, err := os.OpenFile("logs.txt",  os.O_CREATE | os.O_RDWR, 0666)
    //     if err != nil {
    //         fmt.Printf("error opening file: %v", err)
    //     }

    //     // don't forget to close it
    //     defer f.Close()

    //     // assign it to the standard logger
    //     log.SetOutput(f)

	// 	log.Output(1, "this is an event")
		
		
		Log1:=DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000")+" 1st Service"
		ApplicationLogs12 =append(ApplicationLogs12,Log1)
		//fmt.Println(t.Format("2006-01-02 15:04:05.0000")+" 1st Service")
		//fmt.Println(demo2)
		// demo2,_:=fmt.Println("1st Service")
		// fmt.Println(demo1+demo2)
	// var l *Logger
	// l=log.Print(Input)
	//fmt.Println()
	//fmt.Println(Students)
	//fmt.Println(string(reqBody))
	//fmt.Println("here1")
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
//fmt.Println("here2")
	RequestBodyFor2ndService, _ := json.Marshal(Input)
	Log2:=DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000")+string(RequestBodyFor2ndService)
		ApplicationLogs12 =append(ApplicationLogs12,Log2)
	//log.Println(string(RequestBodyFor2ndService))
	Response2ndService, _ := http.Post("http://localhost:10000/SecondService", "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
	
	defer Response2ndService.Body.Close()
	//fmt.Println("here4")
	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	//fmt.Println()
	Log5:=DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000")+"Printing Response"
	ApplicationLogs12 =append(ApplicationLogs12,Log5)
	//log.Println("Printing Response")

	
	Log6:=DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000")+string(ResponseBody)
	ApplicationLogs12 =append(ApplicationLogs12,Log6)
	//log.Println(string(ResponseBody))
	//log.Printf(string(ResponseBody))
	
	fmt.Println()
	//fmt.Println(tim)
	fmt.Println(ApplicationLogs12)
	//fmt.Println(log.Println(string(ResponseBody)))
	s:=string(ResponseBody)
	//fmt.Println("here5")
	json.Unmarshal([]byte(s), &OutputFromSecondService)
	
	json.NewEncoder(w).Encode(OutputFromSecondService)
	fmt.Println("after")
}
func SecondService(w http.ResponseWriter, r *http.Request) {
	DateTimeInMilliseconds := time.Now()
	var Output Output
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)
	
	//fmt.Println()
	Log3:=DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000")+"2nd Service"
		ApplicationLogs12 =append(ApplicationLogs12,Log3)
	//log.Println("2nd Service")
	//fmt.Println()
	//COnverting Input member tpe to string
	out, err := json.Marshal(Input)
	if err != nil {
		panic(err)
	}
	value := string(out)
	Log4:=DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000")+string(value)
	ApplicationLogs12 =append(ApplicationLogs12,Log4)
	//log.Println(Input)
	//fmt.Println()
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
