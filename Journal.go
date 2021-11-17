package main

import (
	"log"
	"os"
	"time"
)

func WriteJournal(journal []string) {
	//Wenn Datei nich existiert Errstelle Diese oder hänge neuen Inhalt an
	file, err := os.OpenFile("Journal/" + time.Now().Format("01-02-2006") + ".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	entry := ""
	for i, s := range journal{
		entry += s
		if(i != len(journal)) {
			entry += ";"
		}
	}
	if _, err:= file.Write([]byte(entry + "\n" )); err != nil {
		file.Close()
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}