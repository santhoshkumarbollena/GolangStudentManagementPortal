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
  
//    myRouter.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
//       clientId := "demo"
//       clientSecret := "demo"
//       err := clientStore.Set(clientId, &models.Client{
//          ID:     clientId,
//          Secret: clientSecret,
//          Domain: "http://localhost:9094",
//       })
//       if err != nil {
//          fmt.Println(err.Error())
//       }

//       w.Header().Set("Content-Type", "application/json")
//       json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
//    })

    //
	

	myRouter.HandleFunc("/FirstService", validateToken(func(w http.ResponseWriter, r *http.Request){
		reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)

	fmt.Println("1st Service")
	fmt.Println()
	fmt.Println(Input)
	fmt.Println()
	//fmt.Println(Students)
	//fmt.Println(string(reqBody))
	RequestBodyFor2ndService, _ := json.Marshal(Input)
	//fmt.Println("------")
	//fmt.Println(r)
	//	fmt.Println("---")

	outd, _ := json.Marshal(r.RequestURI)
	//fmt.Println(string(outd))
	accesskey:=string(outd)
	
	
	dd:=strings.Split(accesskey, "=")
	//fmt.Println(dd[1])
	d:=strings.Split(dd[1], `"`)
	//fmt.Println(d[0])

fmt.Println("--")
fmt.Println("http://localhost:9096/SecondService?access_token="+d[0])
	Response2ndService, _ := http.Post("http://localhost:9096/SecondService?access_token="+d[0], "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
	defer Response2ndService.Body.Close()

	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	fmt.Println()
	fmt.Println("Printing Response")
	fmt.Println(string(ResponseBody))

	json.NewEncoder(w).Encode(string(ResponseBody))
	},srv)).Methods("POST")
	myRouter.HandleFunc("/SecondService", validateToken(func(w http.ResponseWriter, r *http.Request) {
		var Output Output
	Provider := Providers{"demo"}
	var Providers []Providers
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)
	Providers = append(Providers, Provider)
	fmt.Println()
	fmt.Println("2nd Service")
	fmt.Println()
	fmt.Println(Input)
	fmt.Println()
	//Setting output values in service 2
	Output.ResponseId = "responseid"
	Output.Providers = Providers
	Output.NPI = "NPI"
	Output.Ranking = "Ranking"
	Output.ResponseStatus = Object{"ObjectValueValue"}
	Output.StatusCode = "StatusCode"
	Output.StatusMessage = "StatusMessage"

	json.NewEncoder(w).Encode(Output)

	},srv)).Methods("POST")
	
	//

	myRouter.HandleFunc("/protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Hello, I'm protected"))
   }, srv))

   log.Fatal(http.ListenAndServe(":9096", myRouter))
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
	ProviderDemoValue string
}
type Object struct {
	ObjectDemoValue string
}
type Output struct {
	ResponseId     string      `json:"responseId"`
	Providers      []Providers `json:"providers"`
	NPI            string      `json:"NPI"`
	Ranking        string      `json:"ranking"`
	ResponseStatus Object      `json:"responseStatus"`
	StatusCode     string      `json:"statusCode"`
	StatusMessage  string      `json:"statusMessage"`
}
