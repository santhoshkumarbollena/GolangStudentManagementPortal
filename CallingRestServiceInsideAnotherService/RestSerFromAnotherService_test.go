package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
)
func TestFirstService(t *testing.T) {

	var jsonStr = []byte(` {
		"requestId" :                   "dd", 
		"memberId":                     "deo",
		"memberIdType"    :              "dsf",
		"referedToSpecialtyCategory":   "sdfa",
		"referedToSpecialityCode" :      ["demo"],
		"referedToSpecialityAreaOfBody": "sdfsa",
		"providerIds"   :                ["sdfs","asdf"],
		"searchFilterCriteria"   :       "asdfsaf",
		"callingApp"         :       "asdfsa",
		"callingAppType"     :           "asfaf"
	   }`)

	req, err := http.NewRequest("POST", "/FirstService", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FirstService)
	handler.ServeHTTP(rr, req)
//second service test cse
	// req, err = http.NewRequest("POST", "/SecondService", bytes.NewBuffer(jsonStr))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/json")
	// rr = httptest.NewRecorder()
	// handler = http.HandlerFunc(SecondService)
	// handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler status returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"responseId":"1231","providers":[{"NPI":"key1","Ranking":"1"},{"NPI":"key2","Ranking":"2"}],"responseStatus":{"StatusCode":"200","StatusMessage":"sucess"}}`
	
	ss:=string(expected)
	s:=rr.Body.String()
	fmt.Println()
	fmt.Println(s)
	fmt.Println(expected)
	fmt.Println()
	fmt.Println(strings.Compare(strings.TrimRight(s, "\n"), strings.TrimRight(ss, "\n")))
	fmt.Println(strings.Compare(`{"responseId":"1231","providers":[{"NPI":"key1","Ranking":"1"},{"NPI":"key2","Ranking":"2"}],"responseStatus":{"StatusCode":"200","StatusMessage":"sucess"}}`, `{"responseId":"1231","providers":[{"NPI":"key1","Ranking":"1"},{"NPI":"key2","Ranking":"2"}],"responseStatus":{"StatusCode":"200","StatusMessage":"sucess"}}`))
	if strings.Compare(strings.TrimRight(s, "\n"), strings.TrimRight(ss, "\n")) != 0  {
		t.Errorf("handler = returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}