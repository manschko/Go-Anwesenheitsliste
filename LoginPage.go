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

type TemplateDataLogin struct {
	Location string
	Name string
	Key string
	Failed bool
	Success bool

}

var loginData *LoginData = &LoginData{}
var templateDataLogin *TemplateDataLogin = &TemplateDataLogin{"","","",false,false}

func CreateLoginPageServer(port int) *http.Server {



	mux := http.NewServeMux()
	fmt.Println("hit")
	mux.HandleFunc("/", LoginPageHandler)

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func LoginPageHandler( res http.ResponseWriter, req *http.Request) {

	templateDataLogin.Success = false
	//todo check if token exists if not panic hard
	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		templateDataLogin.Failed = true
	}else {
		templateDataLogin.Failed = false
		fmt.Println(keys[0])
		templateDataLogin.Key = keys[0]
	}
	//fmt.Println(keys[0])
	//todo fill TemplateData with Person data if key exists

	tmpl := template.Must(template.ParseFiles( ".\\PageTemplates\\login.html"  ))
	if req.Method != http.MethodPost{
		tmpl.Execute(res, templateDataLogin)
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
		templateDataLogin.Success = true
		tmpl.Execute(res, templateDataLogin)
		fmt.Println(response)
	}

}