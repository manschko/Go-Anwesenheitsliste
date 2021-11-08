package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type LoginData struct {
	Adresse   string
	Name string
}

type TemplateDataLogin struct {
	Location string
	Name string
	Failed bool
	Success bool

}

var loginData *LoginData = &LoginData{}

func CreateLoginPageServer(port int) *http.Server {



	mux := http.NewServeMux()

	mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		data := TemplateDataLogin{"", "", false, true}
		//todo check if token exists if not panic hard
		keys, ok := req.URL.Query()["key"]

		if !ok || len(keys[0]) < 1 {
			data.Failed = true
		}
		//todo fill TemplateData with Person data if key exists

		tmpl := template.Must(template.ParseFiles( ".\\PageTemplates\\login.html"  ))
		if req.Method != http.MethodPost{
			tmpl.Execute(res, data)
			return
		}


		req.ParseForm()
		loginData.Name = req.FormValue("name")
		loginData.Adresse = req.FormValue("adresse")
		response := LoginData{
			Name:   req.FormValue("name"),
			Adresse: req.FormValue("adresse"),
		}
		if(response  != LoginData{}) {
			//todo get location from keys and login status form list
			WriteJournal(response.Name, "test", true)
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