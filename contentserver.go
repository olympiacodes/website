package main

import (
	"html/template"
	"net/http"
)

// HTMLServer is a http.Handler for serving HTML content in the site theme. It
// is primarily designed for use with simple HTML content, generally converted
// from Markdown documents.
type HTMLServer struct {
	GroupName string
	Title     string
	Content   template.HTML
}

// ServeHTTP implements the http.Handler interface, serving the Content HTML
// in the site theme.
func (h *HTMLServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	templateData, err := Asset("content.html")
	if err != nil {
		rw.Write([]byte("Error processing request, please try again later."))
		return
	}

	t, err := template.New("homepage").Parse(string(templateData))
	if err != nil {
		rw.Write([]byte("Error processing request, please try again later."))
		return
	}

	t.Execute(rw, h)
}
