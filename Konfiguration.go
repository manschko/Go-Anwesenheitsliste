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
var wg = new(sync.WaitGroup)
func ConfigFlag() {
	flag.IntVar(&flags.Port1, "portLogin", 8000, "HTTP Server port")
	flag.IntVar(&flags.Port2, "portQR", 8080, "HTTP Server port")
	flag.IntVar(&flags.TokenValidity, "valid", 5, "Gültigkeitsdauer der Token")

	flag.Parse()

}

func ConfigWebServer(){
	//WaitGroup für go routinen erstellt
	//Setzte WaitGroup auf 2 für 2 go routinen
	wg.Add(2)

	// Setup für Anmeldeserver über go routine
	go func() {
		server := CreateLoginPageServer("Login", flags.Port1)
		fmt.Println( server.ListenAndServe())
		wg.Done()
	}()

	// Setup für QRcode Seite über go routinetes
	go func() {
		server := createServer("QR", flags.Port2)
		fmt.Println( server.ListenAndServe())
		wg.Done()
	}()


}

func createServer(name string, port int)  *http.Server{

	mux := http.NewServeMux()

	mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: " + name)
	})

	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}





