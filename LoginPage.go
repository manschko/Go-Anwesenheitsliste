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

var loginData *LoginData = &LoginData{}

func CreateLoginPageServer(name string, port int)  *http.Server{

	mux := http.NewServeMux()

	mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {

		renderTemplate(res, req)
	})

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}


func renderTemplate(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles( ".\\PageTemplates\\login.html"))

	if req.Method != http.MethodPost{
		tmpl.Execute(res, nil)
		return
	}

	loginData.Name = req.FormValue("name")
	loginData.Adresse = req.FormValue("adresse")
	response := LoginData{
		Name:   req.FormValue("name"),
		Adresse: req.FormValue("adresse"),
	}
	//todo call Journal
	fmt.Println(response)

	tmpl.Execute(res, struct{ Success bool}{true})
}