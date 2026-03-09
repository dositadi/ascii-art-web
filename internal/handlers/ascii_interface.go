package handlers

import (
	m "ascii-web/pkg/models"
	"log"
	"net/http"
)

type AsciiInterface interface {
	Get() *m.Error
	Post(request m.Ascii) (m.Response, *m.Error)
}

type AsciiSketch struct {
	AsciiSketcher AsciiInterface
}

func CreateNewAsciiSketcher(sketcher AsciiInterface) *AsciiSketch {
	return &AsciiSketch{
		AsciiSketcher: sketcher,
	}
}

func (s *AsciiSketch) HandleVerbs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.PostHandler(w, r)
	case "GET":
		s.GetHandler(w, r)
	}
}

func (s *AsciiSketch) GetHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *AsciiSketch) PostHandler(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("text")
	banner := r.PostFormValue("banner")

	request := m.Ascii{
		Text:   text,
		Banner: banner,
	}

	response, err := s.AsciiSketcher.Post(request)
	if err != nil {
		log.Printf(err.Error, ": ", err.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response.Reponse))
}
