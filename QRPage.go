package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
	qrcode "github.com/skip2/go-qrcode"
	"strconv"
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

		m.HandleFunc("/{location}", func( res http.ResponseWriter, req *http.Request) {
			//data := TemplateData{[]string{"test","test2"}}

			params := mux.Vars(req)

			for _, location := range locations {

				if params["location"] == location.Name {

					//TODO vergleiche ob params in locations, wenn nicht vorhanden startseite; return , wenn vorhanden ausführen
					path := filepath.FromSlash("./PageTemplates/qrSingle.html")
					tmpl := template.Must(template.ParseFiles(path))

					fmt.Println(location.Name)

					if req.Method != http.MethodPost {
						executeQr(location)
						tmpl.Execute(res, struct{ locationName string }{location.Name})
						return
					}

					//Todo if form got send

					fmt.Println(req.FormValue("location"))
					//renderTemplate(res, req, "qr.html", TemplateData{[]string{}, true})

				} else {
					//TODO startseite laden
				return
				}
			}
		})
	}

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: m,
	}

	return &server
}

func executeQr(location Location) {


	println(location.AccessToken + " Access")
	println(location.CurrentToken + " Current")

	qrcode.WriteFile("https://localhost:" + strconv.Itoa(flags.Port1) + "/?location=" + location.AccessToken + "&access=" + location.CurrentToken, qrcode.Medium, 256, "PageTemplates/qr.png")

}
