package webapp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"thinknetica/Lesson-12/pkg/crawler"
	"thinknetica/Lesson-12/pkg/webapp/response"
)

type Controller struct {
	index       map[string][]int
	scanResults []crawler.Document
}

func NewController(index map[string][]int, scanResults []crawler.Document) *Controller {
	return &Controller{index: index, scanResults: scanResults}
}

func (c *Controller) ShowIndexData(w http.ResponseWriter, r *http.Request) {
	format := getResponseFormat(r)
	indexData := createIndexData(c.index, format)

	switch format {
	case "json":
		writeJSONResponse(w, indexData)
	case "xml":
		writeXMLResponse(w, indexData)
	case "html":
		writeHTMLResponse(w, indexData, "index")
	default:
		http.Error(w, "Unsupported format", http.StatusBadRequest)
	}
}

func (c *Controller) ShowDocData(w http.ResponseWriter, r *http.Request) {
	format := getResponseFormat(r)
	docData := createDocData(c.scanResults, format)

	switch format {
	case "json":
		writeJSONResponse(w, docData)
	case "xml":
		writeXMLResponse(w, docData)
	case "html":
		writeHTMLResponse(w, docData, "doc")
	default:
		http.Error(w, "Unsupported format", http.StatusBadRequest)
	}
}

func (c *Controller) ShowIndexDataXML(w http.ResponseWriter, r *http.Request) {
	indexData := createIndexData(c.index, "xml")
	writeXMLResponse(w, indexData)
}

func (c *Controller) ShowIndexDataHTML(w http.ResponseWriter, r *http.Request) {
	indexData := createIndexData(c.index, "html")
	writeHTMLResponse(w, indexData, "index")
}

func getResponseFormat(r *http.Request) string {
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "json"
	}
	return format
}

func createIndexData(index map[string][]int, format string) interface{} {
	switch format {
	case "json":
		var indexData []*response.JSONIndexData
		for key, val := range index {
			indexData = append(indexData, &response.JSONIndexData{Token: key, PositionList: fmt.Sprintf("%v", val)})
		}
		return indexData
	case "xml":
		var indexData []*response.XMLIndexData
		for key, val := range index {
			indexData = append(indexData, &response.XMLIndexData{Token: key, PositionList: fmt.Sprintf("%v", val)})
		}
		return indexData
	case "html":
		var indexData string
		for key, val := range index {
			indexData += fmt.Sprintf("<p>%s: %v</p>", key, val)
		}
		return indexData
	default:
		return nil
	}
}

func createDocData(scanResults []crawler.Document, format string) interface{} {
	switch format {
	case "json":
		var docData []*response.DocData
		for _, val := range scanResults {
			docData = append(docData, &response.DocData{Title: val.Title, Body: val.Body, URL: val.URL})
		}
		return docData
	case "xml":
		var docData []*response.XMLDocData
		for _, val := range scanResults {
			docData = append(docData, &response.XMLDocData{Title: val.Title, Body: val.Body, URL: val.URL})
		}
		return docData
	case "html":
		var docData string
		for _, val := range scanResults {
			docData += fmt.Sprintf("<h2>%s</h2><p>%s</p><a href=\"%s\">Link</a><br>", val.Title, val.Body, val.URL)
		}
		return docData
	default:
		return nil
	}
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func writeXMLResponse(w http.ResponseWriter, data interface{}) {
	result, err := xml.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding XML", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func writeHTMLResponse(w http.ResponseWriter, data interface{}, templateName string) {
	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{.Data}}
	</body>
	</html>
	`

	tmpl, err := template.New("htmlTemplate").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	title := templateName + " data"
	pageData := struct {
		Title string
		Data  interface{}
	}{
		Title: title,
		Data:  data,
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Error rendering HTML", http.StatusInternalServerError)
		return
	}
}
