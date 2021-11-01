package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)
//Struct für das Speichern der Flags
type Flags struct {
	Port1 int
	Port2 int
	TokenValidity int
}
//var um global auf die Flags zugreifen zu können
var flags *Flags = &Flags{}

func ConfigFlag() {
	flag.IntVar(&flags.Port1, "portLogin", 8000, "HTTP Server port")
	flag.IntVar(&flags.Port2, "portQR", 8080, "HTTP Server port")
	flag.IntVar(&flags.TokenValidity, "valid", 5, "Gültigkeits dauer der Token")

	flag.Parse()

}

func ConfigWebServer(){
	//WaitGroup für go routinen erstellt
	wg := new(sync.WaitGroup)
	//Setzte WaitGroup auf 2 für 2 go routinen
	wg.Add(2)

	// Setup für Anmeldeserver über go routine
	go func() {
		server := createServer("Login", flags.Port1)
		fmt.Println(server.ListenAndServeTLS("cert.pem", "key.pem"))
		wg.Done()
	}()

	// Setup für QRcode Seite über go routine
	go func() {
		server := createServer("QR", flags.Port2)
		fmt.Println(server.ListenAndServeTLS("cert.pem", "key.pem"))
		wg.Done()
	}()

	wg.Wait()
}

func createServer(name string, port int)  *http.Server{

	mux := http.NewServeMux()

	// Ursprüngliche Ausgabe Inhalt
	/*mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: " + name)
	})*/

	//Auswahl der HTML Seite je nach Port
	// Sollte angepasst werden, habe es nur mit jeweils einer index.html
	//geschafft, deshalb die verschiedenen Ordner
	if (name == "Login") {
		mux.Handle("/", http.FileServer(http.Dir("./static/login")))
	}
	if(name == "QR") {
		mux.Handle("/", http.FileServer(http.Dir("./static/qr")))
	}

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}