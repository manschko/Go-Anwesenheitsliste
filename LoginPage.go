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
	Email   string
	Subject string
	Message string
}

var loginData *LoginData = &LoginData{}

func CreateLoginPageServer(name string, port int)  *http.Server{

	mux := http.NewServeMux()

	mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {

		renderTemplate(res, req)
	})

	mux.HandleFunc("/result", func( res http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		for key, value := range req.Form {
			fmt.Printf("%s = %s\n", key, value)
		}
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

	loginData.Email = req.FormValue("email")
	loginData.Subject = req.FormValue("subject")
	loginData.Message = req.FormValue("message")
	response := LoginData{
		Email:   req.FormValue("email"),
		Subject: req.FormValue("subject"),
		Message: req.FormValue("message"),
	}
	//todo call Journal
	fmt.Println(response)

	tmpl.Execute(res, struct{ Success bool}{true})
}