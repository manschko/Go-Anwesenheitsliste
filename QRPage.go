package main

import (
	"fmt"
	"github.com/gorilla/mux"
	qrcode "github.com/skip2/go-qrcode"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type TemplateDataQR struct {
	Locations []string
	Success bool
}

func createQRWebServer(port int)  *http.Server{
	m := mux.NewRouter()
	locations, check := ReadLocationList()

	//show form with locations for selection
	m.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		path := filepath.FromSlash("./PageTemplates/qr.html")
		tmpl := template.Must(template.ParseFiles(path))
		if req.Method != http.MethodPost{
			//TODO inhalt aus xml hinzufügen als auswahlmöglichkeiten
			tmpl.Execute(res, TemplateDataQR{[]string{"Mosbach", "test2"},false})
			return
		}
		//TODO eventuell abfangen wenn falsches ausgewählt wird das nicht gibt
		//get selected location
		selectedLocation := req.FormValue("location")

		for _, location := range locations {
			fmt.Println(location)
			//check if selected location is in xml
			if selectedLocation == location.Name {
				//add accesstoken of selected location to url
				http.Redirect(res, req, "/" + location.AccessToken, http.StatusSeeOther)
			}
		}
	})
	if check {
		//get location as param from url
		m.HandleFunc("/{location}", func( res http.ResponseWriter, req *http.Request) {
			params := mux.Vars(req)
			for _, location := range locations {
				//check if param is in xml
				if params["location"] == location.AccessToken {
					//load new template for qr code
					path := filepath.FromSlash("./PageTemplates/qrSingle.html")
					tmpl := template.Must(template.ParseFiles(path))
					//generate qr code
					executeQr(location)
					tmpl.Execute(res, struct{LocationName string
											Valid int}{location.Name, flags.TokenValidity})
					return

					// if param is not in xml redirect back to selection page
				} else {
					http.Redirect(res, req, "/", http.StatusSeeOther)
					return
				}
				return
			}
		})
	}
	//fileserver for gr code single page
	fs := http.FileServer( http.Dir("./PageTemplates"))
	m.PathPrefix("/{location}").Handler(http.StripPrefix("/PageTemplates", fs))
	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: m,
	}
	return &server
}
//generate qr code as png and save it in PageTemplates
func executeQr(location Location) {

	qrcode.WriteFile("https://localhost:" + strconv.Itoa(flags.Port1) + "/?location=" + location.AccessToken + "&access=" + location.CurrentToken, qrcode.Medium, 256, "PageTemplates/qr.png")
}
