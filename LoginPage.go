package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

type LoginData struct {
	Adresse   string
	Name string
}

type TemplateData struct {
	Locations []string
	Success bool
}

var loginData *LoginData = &LoginData{}

func CreateLoginPageServer(name string, port int) *http.Server {

	//todo check if token exists if not panic hard
	mux := http.NewServeMux()

	mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		//todo fill TemplateData with Person data if key exists
		data := TemplateData{nil, false}
		tmpl := template.Must(template.ParseFiles( ".\\PageTemplates\\login.html"  ))
		if req.Method != http.MethodPost{
			tmpl.Execute(res, data)
			return
		}

		//todo call Journal
		req.ParseForm()
		loginData.Name = req.FormValue("name")
		loginData.Adresse = req.FormValue("adresse")
		response := LoginData{
			Name:   req.FormValue("name"),
			Adresse: req.FormValue("adresse"),
		}
		if(response  != LoginData{}) {
			data.Success = true
			tmpl.Execute(res, data)
			fmt.Println(response)
		}

	})

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}