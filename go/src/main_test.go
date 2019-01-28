//main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"log"
	"github.com/gorilla/mux"
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

func TestRouter (t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", r))

	server := httptest.NewServer(r)

	resp,err := http.Get(server.URL + "/users")
	if err != nil {t.Fatal(err)}

	//if status code is not 200 print error
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Got bad status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {t.Fatal(err)}

	responseString := string(body)

	if responseString != "userHandler hit!" {
		t.Errorf("Unexpected response : %s", responseString)
	}
}