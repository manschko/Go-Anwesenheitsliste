package main

import (
	"flag"
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



