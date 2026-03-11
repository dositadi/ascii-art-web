package handlers

import (
	m "ascii-web/pkg/models"
	h "ascii-web/pkg/utils"
	"fmt"
	"net/http"
	"text/template"
)

// This is the interface that contains the methods the its type mus implement
type AsciiInterface interface {
	Post(request m.Ascii) (string, *m.Error)
}

// This struct only acept a field type that implements the AsciiInterface 
type AsciiSketch struct {
	AsciiSketcher AsciiInterface
}

// A contructor to create a new Asci sketch
func CreateNewAsciiSketcher(sketcher AsciiInterface) *AsciiSketch {
	return &AsciiSketch{
		AsciiSketcher: sketcher,
	}
}

// This is a wrapper handler that checks for the nmethod used in a request and respond respectively
func (s *AsciiSketch) HandleVerbs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.PostHandler(w, r)
	case "GET":
		s.GetHandler(w, r)
	}
}

// This handler handles get request
func (s *AsciiSketch) GetHandler(w http.ResponseWriter, r *http.Request) {
	// Template initialization for the home page
	temp, err := template.New("index.html").ParseFiles("web/index.html")
	if err != nil {
		errMessage := fmt.Sprintf("Error: %s\nDetail: %s\n", h.SERVER_ERR, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errMessage))
		return
	}

	pageDetails := m.ASCIIWeb{Title: "Ascii-Art Web App", PostUrl: "http://localhost:8081/ascii-art", HomeUrl: "http://localhost:8081/"}

	err2 := temp.Execute(w, pageDetails)
	if err2 != nil {
		errMessage := fmt.Sprintf("Error: %s\nDetail: %s\n", h.SERVER_ERR, err2.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errMessage))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}

// this handler handles post request
func (s *AsciiSketch) PostHandler(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("text")
	banner := r.PostFormValue("banner")

	if text == "" || banner == "" {
		errorMessage := "Error: Invalid Input"
		h.Response(w, r, http.StatusBadRequest, errorMessage)
		return
	}

	request := m.Ascii{
		Text:   text,
		Banner: banner,
	}

	response, err := s.AsciiSketcher.Post(request)
	if err != nil {
		errMessage := fmt.Sprintf("Error: %s\nDetail: %s\n", err.Error, err.Detail)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errMessage))
	}

	responsePage := struct {
		Title    string
		HomeUrl  string
		Input    string
		Response string
	}{
		Title:    "Ascii-Art Web App",
		HomeUrl:  "http://localhost:8081/",
		Input:    text,
		Response: response,
	}

	temp, err2 := template.New("response.html").ParseFiles("web/response.html")
	if err != nil {
		errMessage := fmt.Sprintf("Error: %s\nDetail: %s\n", h.SERVER_ERR, err2.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errMessage))
		return
	}

	err3 := temp.Execute(w, responsePage)
	if err3 != nil {
		errMessage := fmt.Sprintf("Error: %s\nDetail: %s\n", h.SERVER_ERR, err3.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errMessage))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/html")
}
