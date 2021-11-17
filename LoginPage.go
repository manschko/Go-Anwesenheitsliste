package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type Page struct{
	Title string
	Body []byte
}

type LoginData struct {
	Adresse string
	Name    string
	Login   bool
}

type TemplateDataLogin struct {
	LocationToken  string
	Location       string
	Adresse        string
	Name           string
	Failed         bool
	Success        bool
}

var dataMap map[string] TemplateDataLogin  = make(map[string]TemplateDataLogin )
var loginData *LoginData = &LoginData{}
var templateDataLogin *TemplateDataLogin = &TemplateDataLogin{"", "","","",false,false}

func CreateLoginPageServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", LoginPageHandler)

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func LoginPageHandler( res http.ResponseWriter, req *http.Request) {

	templateDataLogin.Success = false
	location := req.URL.Query().Get("location")
	access := req.URL.Query().Get("access")

	if location == "" && access == "" {
		templateDataLogin.Failed = true
	}else {
		templateDataLogin.Failed = false
		templateDataLogin.Name = ""
		templateDataLogin.Location = ""
		templateDataLogin.LocationToken = ""
		templateDataLogin.Adresse = ""
		updateTokenList(access, location)
	}
	path := filepath.FromSlash("./PageTemplates/login.html")
	tmpl := template.Must(template.ParseFiles(path))
	if req.Method != http.MethodPost{
		tmpl.Execute(res, templateDataLogin)
		return
	}


	req.ParseForm()
	loginData.Name = req.FormValue("name")
	loginData.Adresse = req.FormValue("adresse")

	loginData.Login, _ = strconv.ParseBool(req.FormValue("submit"))
	var response = LoginData{
		Name:    loginData.Name,
		Adresse: loginData.Adresse,
		Login:   loginData.Login,
	}
	if(response  != LoginData{}) {
		templateDataLogin.Name = response.Name
		templateDataLogin.Adresse = response.Adresse
		if(templateDataLogin.Location != ""){
			entry := []string{templateDataLogin.Location, templateDataLogin.Adresse, templateDataLogin.Name}
			if loginData.Login {
				entry = append(entry, "Angemeldet")
			}else {
				entry = append(entry,"Abgemeldet")
			}
			WriteJournal(entry)
			templateDataLogin.Success = true
		}else{
			templateDataLogin.Failed = true
		}
		dataMap[access] = *templateDataLogin

		if(!response.Login) {
			delete(dataMap, access)
		}
		tmpl.Execute(res, templateDataLogin)
	}

}

func updateTokenList(access string, location string) {

	for key, value := range dataMap{
		if !isTokenValid(value.LocationToken, key) {
			delete(dataMap, access)
		}
	}

	_, ok := dataMap[access]

	if !ok {
		locations, _:= ReadLocationList()
		for i := range locations{
			if locations[i].AccessToken == location {
				templateDataLogin.Location 	= locations[i].Name
			}
		}
		templateDataLogin.LocationToken = location
		dataMap[access] = *templateDataLogin
		return
	}
	*templateDataLogin = dataMap[access]
	templateDataLogin.Success = false
}
