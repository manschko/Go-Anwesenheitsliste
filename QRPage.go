package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"fmt"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type TemplateDataQR struct {
	Locations []string
	Success   bool
}

var locations, check = ReadLocationList()
var locationNameList []string

func createQRWebServer(port int) *http.Server {
	m := http.NewServeMux()

	// list of location names

	for _, location := range locations {
		locationNameList = append(locationNameList, location.Name)
	}

	//show form with locations for selection
	m.HandleFunc("/", selectionPageHandler)
	//fileserver for gr code single page
	fs := http.FileServer(http.Dir("./PageTemplates"))
	m.Handle("/assets/", http.StripPrefix("/assets/", fs))
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: m,
	}
	return &server
}

//generate qr code as png and save it in PageTemplates
func executeQr(location Location) {

	qrcode.WriteFile("https://"+flags.Url+":"+strconv.Itoa(flags.Port1)+"/?location="+location.AccessToken+"&access="+location.CurrentToken, qrcode.Medium, 256, "PageTemplates/qr.png")
}
func qrPageHandler(res http.ResponseWriter, req *http.Request, param string) {

	for _, location := range locations {
		//check if param is in xml
		if param == location.AccessToken {
			//load new template for qr code
			path := filepath.FromSlash("./PageTemplates/qrSingle.html")
			tmpl := template.Must(template.ParseFiles(path))
			//generate qr code
			executeQr(location)
			tmpl.Execute(res, struct {
				LocationName string
				Valid        int
			}{location.Name, flags.TokenValidity})
			return
		}
	}
	// if param is not in xml redirect back to selection page
	http.Redirect(res, req, "/", http.StatusSeeOther)
	return
}

func selectionPageHandler(res http.ResponseWriter, req *http.Request) {
	splittedUrl := strings.Split(req.URL.Path[1:], "/")
	location := splittedUrl[len(splittedUrl)-1]
	if location == "" {

		path := filepath.FromSlash("./PageTemplates/qr.html")
		tmpl := template.Must(template.ParseFiles(path))
		if req.Method != http.MethodPost {
			tmpl.Execute(res, TemplateDataQR{locationNameList, false})
			return
		}
		//get selected location
		selectedLocation := req.FormValue("location")

		for _, location := range locations {
			//check if selected location is in xml
			if selectedLocation == location.Name {
				//add accesstoken of selected location to url
				http.Redirect(res, req, "/"+location.AccessToken, http.StatusSeeOther)
			}
		}
	} else {
		qrPageHandler(res, req, location)
	}
}
