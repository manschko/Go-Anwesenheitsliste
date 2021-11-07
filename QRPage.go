package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func createQRWebServer(name string, port int)  *http.Server{
	mux := http.NewServeMux()

	// Ursprüngliche Ausgabe Inhalt
	mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		//data := TemplateData{[]string{"test","test2"}}
		tmpl := template.Must(template.ParseFiles( ".\\PageTemplates\\qr.html"  ))
		if req.Method != http.MethodPost{
			tmpl.Execute(res, TemplateData{[]string{"test", "test2"},false})
			return
		}
		//Todo if form got send
		fmt.Println(req.FormValue("location"))
		//renderTemplate(res, req, "qr.html", TemplateData{[]string{}, true})
	})

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}