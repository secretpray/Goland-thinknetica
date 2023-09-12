package webapp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"thinknetica/Lesson-12/pkg/crawler"
)

func TestShowIndexData(t *testing.T) {
	expected := `[{"token":"Golang","positions_list":"[1 2 3]"},{"token":"Paypal","positions_list":"[4 5 6]"}][{"token":"Golang","positions_list":"[1 2 3]"},{"token":"Paypal","positions_list":"[4 5 6]"}]`
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
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
	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}

func TestShowDocData(t *testing.T) {
	expected := `[{"title":"go-dev","body":"","url":"https://go.dev/"},{"title":"golang org","body":"","url":"https://golang-org.appspot.com/"}][{"title":"go-dev","body":"","url":"https://go.dev/"},{"title":"golang org","body":"","url":"https://golang-org.appspot.com/"}]`
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
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
	if string(responseBody) != expected {
		t.Errorf("Expected %s but got %v", expected, string(responseBody))
	}
}
