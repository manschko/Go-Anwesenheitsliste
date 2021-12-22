package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"flag"
	"fmt"
	"sync"
)

//Struct für das Speichern der Flags
type Flags struct {
	Port1         int
	Port2         int
	TokenValidity int
	Url           string
}

//var um global auf die Flags zugreifen zu können
var flags *Flags = &Flags{}
var wg = new(sync.WaitGroup)

func ConfigFlag() {
	flag.IntVar(&flags.Port1, "portLogin", 8000, "HTTP Server port")
	flag.IntVar(&flags.Port2, "portQR", 8080, "HTTP Server port")
	flag.IntVar(&flags.TokenValidity, "valid", 3600, "Gültigkeitsdauer der Token")
	flag.StringVar(&flags.Url, "url", "localhost", "URL für den Webserver")
	flag.Parse()

}

func ConfigWebServer() {
	//WaitGroup für go routinen erstellt
	//Setzte WaitGroup auf 2 für 2 go routinen
	wg.Add(2)

	// Setup für Anmeldeserver über go routine
	go func() {
		server := CreateLoginPageServer(flags.Port1)
		fmt.Println(server.ListenAndServeTLS("cert.pem", "key.pem"))
		wg.Done()
	}()

	// Setup für QRcode Seite über go routinetes
	go func() {
		server := createQRWebServer(flags.Port2)
		fmt.Println(server.ListenAndServeTLS("cert.pem", "key.pem"))
		wg.Done()
	}()

}
