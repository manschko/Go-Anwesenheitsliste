package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func WriteJournal(name string, ort string, anmelden bool ) {
	//Wenn Datei nich existiert Errstelle Diese oder hänge neuen Inhalt an
	file, err := os.OpenFile("Journal/" + time.Now().Format("01-02-2006"), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	if _, err:= file.Write([]byte(name + ";" + ort + ";" + strconv.FormatBool(anmelden) + "\n" )); err != nil {
		file.Close()
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}