package main

import (
	
	"fmt"

	"net/http"



	"github.com/gorilla/mux"
)

var host = "http://localhost"
var port = "12345"
var connectionString = "root:7798775575@tcp(127.0.0.1:3306)/online_address_book?charset=utf8&parseTime=True&loc=Local"

func main() {

	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()                                               // /api will give access to all the API endpoints
	apiRouter.PathPrefix("/entries").HandlerFunc(GetEntries).Methods("GET")                          // /api/entries returns listing all the entries
						  // DELETE /api/entry deletes an entry
						  	fmt.Println("Listening on port :12345")
	http.ListenAndServe(":"+port, router)

}

// GetEntries : Get All Entries
// URL : /entries
// Parameters: none
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
func GetEntries(w http.ResponseWriter, r *http.Request) {
	entries:=`[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"},{"id":2,"first_name":"Rick","last_name":"Parker","email_address":"rick.parker@gmail.com","phone_number":"1234567890"},{"id":3,"first_name":"Kelly","last_name":"Franco","email_address":"kelly.franco@gmail.com","phone_number":"1112223333"}]`
	respondWithJSON(w, http.StatusOK, entries)
}