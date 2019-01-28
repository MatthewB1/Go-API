//main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//create test requests (3 args)
	req, err := http.NewRequest("GET", "",  nil)
	if err != nil {t.Fatal(err)}


	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(userHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned bad status code: %v", status)
	}

	returnedBody := recorder.Body.String()
	if returnedBody != "userHandler hit!" {
		t.Errorf("handler returned unexpected body : %v", returnedBody)
	}

	// fmt.Printf("test(s) ran succesfully")
	// os.Exit(0)
}