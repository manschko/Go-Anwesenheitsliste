package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateDataQR struct {
	Locations []string
	Success bool
}

func createQRWebServer(port int)  *http.Server{
	m := mux.NewRouter()

	m.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		//data := TemplateData{[]string{"test","test2"}}
		path := filepath.FromSlash("./PageTemplates/qr.html")
		tmpl := template.Must(template.ParseFiles(path))
		if req.Method != http.MethodPost{
			tmpl.Execute(res, TemplateDataQR{[]string{"test", "test2"},false})
			return
		}
		//Todo if form got send
		//getNewUrl(res, req)

		fmt.Println(req.FormValue("location"))
		//renderTemplate(res, req, "qr.html", TemplateData{[]string{}, true})
	})

	locations, check := ReadLocationList()
	if check {
		for _, location := range locations {

			m.HandleFunc("/{location}", func( res http.ResponseWriter, req *http.Request) {
				//data := TemplateData{[]string{"test","test2"}}
				params := mux.Vars(req)
				_ = params
				path := filepath.FromSlash("./PageTemplates/qrSingle.html")
				tmpl := template.Must(template.ParseFiles(path))

				fmt.Println(location.Name)

				if req.Method != http.MethodPost{
					tmpl.Execute(res, struct{locationName string}{location.Name})
					return
				}
				//Todo if form got send

				fmt.Println(req.FormValue("location"))
				//renderTemplate(res, req, "qr.html", TemplateData{[]string{}, true})
			})

		}
	}

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: m,
	}

	return &server
}

