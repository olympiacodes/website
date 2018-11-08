package main

import (
	"html/template"
	"net/http"

	"github.com/olympiacodes/website/internal/meetup"
)

type HomePageServer struct {
	fileServer http.Handler

	GroupName string
	Events    []meetup.Event
	Images    []Image

	TwitterUsername   string
	InstagramUsername string
	FacebookPage      string
	MeetupGroupName   string
}

type Image struct {
	Src  string
	Link string
	Alt  string
}

func (s *HomePageServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Index Page
	if req.URL.Path == "/" {
		s.handleIndex(rw, req)
		return
	}

	if s.fileServer == nil {
		s.fileServer = http.FileServer(Assets)
	}

	s.fileServer.ServeHTTP(rw, req)
}

func (s *HomePageServer) handleIndex(rw http.ResponseWriter, req *http.Request) {
	templateData, err := Asset("index.html")
	if err != nil {
		rw.Write([]byte("Error processing request, please try again later."))
		return
	}

	t, err := template.New("homepage").Parse(string(templateData))
	if err != nil {
		rw.Write([]byte("Error processing request, please try again later."))
		return
	}

	t.Execute(rw, s)
}
