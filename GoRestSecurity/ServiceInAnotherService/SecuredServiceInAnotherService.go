package main

import (
   "encoding/json"
   "fmt"
    "strings"
   "gopkg.in/oauth2.v3/models"
   "log"
   "net/http"
   "io/ioutil"
   "bytes"
   "gopkg.in/oauth2.v3/errors"
   "gopkg.in/oauth2.v3/manage"
   "gopkg.in/oauth2.v3/server"
   "github.com/gorilla/mux"
   "gopkg.in/oauth2.v3/store"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
   manager := manage.NewDefaultManager()
   manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

   // token memory store
   manager.MustTokenStorage(store.NewMemoryTokenStore())

   // client memory store
   clientStore := store.NewClientStore()
   
   manager.MapClientStorage(clientStore)

   srv := server.NewDefaultServer(manager)
   srv.SetAllowGetAccessRequest(true)
   srv.SetClientInfoHandler(server.ClientFormHandler)
   manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

   srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
      log.Println("Internal Error:", err.Error())
      return
   })

   srv.SetResponseErrorHandler(func(re *errors.Response) {
      log.Println("Response Error:", re.Error.Error())
   })

   myRouter.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
	clientId := "demo"
	clientSecret := "demo"
	err := clientStore.Set(clientId, &models.Client{
	   ID:     clientId,
	   Secret: clientSecret,
	   Domain: "http://localhost:9094",
	})
	if err != nil {
	   fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
      srv.HandleTokenRequest(w, r)
   })
  

	

	myRouter.HandleFunc("/FirstService", validateToken(func(w http.ResponseWriter, r *http.Request){
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
	

	ToGetAccessKeyFromFirstService, _ := json.Marshal(r.RequestURI)

	BodyContainingAccesskey:=string(ToGetAccessKeyFromFirstService)
	
	
	AccessKeyWithDoubleQuoteAtEnd:=strings.Split(BodyContainingAccesskey, "=")

	AccessKeyToPassToSecondService:=strings.Split(AccessKeyWithDoubleQuoteAtEnd[1], `"`)
	
	Response2ndService, _ := http.Post("http://localhost:10000/SecondService?access_token="+AccessKeyToPassToSecondService[0], "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
	defer Response2ndService.Body.Close()

	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	fmt.Println()
	fmt.Println("Printing Response")
	fmt.Println(string(ResponseBody))
	s:=string(ResponseBody)
	
    json.Unmarshal([]byte(s), &OutputFromSecondService)
	json.NewEncoder(w).Encode(OutputFromSecondService)
	},srv)).Methods("POST")


	myRouter.HandleFunc("/SecondService", validateToken(func(w http.ResponseWriter, r *http.Request) {
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

	},srv)).Methods("POST")
	


	

   log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  _, err := srv.ValidationBearerToken(r)
	 
      if err != nil {
         http.Error(w, err.Error(), http.StatusBadRequest)
         return
      }

      f.ServeHTTP(w, r)
   })
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
