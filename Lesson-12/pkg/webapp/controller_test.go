package webapp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"thinknetica/Lesson-12/pkg/crawler"
)

func TestShowIndexData(t *testing.T) {
	expectedJSON := `[{"token":"Golang","positions_list":"[1 2 3]"},{"token":"Paypal","positions_list":"[4 5 6]"}]`
	req := httptest.NewRequest(http.MethodGet, "/index?format=json", nil)
	w := httptest.NewRecorder()
	var index = map[string][]int{"Golang": {1, 2, 3}, "Paypal": {4, 5, 6}}
	c := NewController(index, make([]crawler.Document, 0))
	c.ShowIndexData(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expectedJSON {
		t.Errorf("Expected %s but got %v", expectedJSON, string(data))
	}

	// Тестирование формата XML
	expectedXML := `<XMLIndexData><token>Golang</token><positions_list>[1 2 3]</positions_list></XMLIndexData><XMLIndexData><token>Paypal</token><positions_list>[4 5 6]</positions_list></XMLIndexData>`
	reqXML := httptest.NewRequest(http.MethodGet, "/index?format=xml", nil)
	wXML := httptest.NewRecorder()
	c.ShowIndexData(wXML, reqXML)
	resXML := wXML.Result()
	defer resXML.Body.Close()
	dataXML, err := io.ReadAll(resXML.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(dataXML) != expectedXML {
		t.Errorf("Expected XML:\n%s\nbut got:\n%v", expectedXML, string(dataXML))
	}
}

func TestShowDocData(t *testing.T) {
	expectedJSON := `[{"title":"go-dev","body":"","url":"https://go.dev/"},{"title":"golang org","body":"","url":"https://golang-org.appspot.com/"}]`
	req := httptest.NewRequest(http.MethodGet, "/doc?format=json", nil)
	w := httptest.NewRecorder()
	var index = map[string][]int{"Golang": {1, 2, 3}, "Paypal": {4, 5, 6}}
	data := []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev/",
			Title: "go-dev",
		},
		{
			ID:    1,
			URL:   "https://golang-org.appspot.com/",
			Title: "golang org",
		},
	}
	c := NewController(index, data)
	c.ShowDocData(w, req)
	res := w.Result()
	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(responseBody) != expectedJSON {
		t.Errorf("Expected %s but got %v", expectedJSON, string(responseBody))
	}

	// Тестирование формата XML
	expectedXML := `<XMLDocData><Title>go-dev</Title><Body></Body><URL>https://go.dev/</URL></XMLDocData><XMLDocData><Title>golang org</Title><Body></Body><URL>https://golang-org.appspot.com/</URL></XMLDocData>`

	reqXML := httptest.NewRequest(http.MethodGet, "/doc?format=xml", nil)
	wXML := httptest.NewRecorder()
	c.ShowDocData(wXML, reqXML)
	resXML := wXML.Result()
	defer resXML.Body.Close()
	dataXML, err := io.ReadAll(resXML.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(dataXML) != expectedXML {
		t.Errorf("Expected XML:\n%s\nbut got:\n%v", expectedXML, string(dataXML))
	}
}
