package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"errors"
	"log"
	"os"
	"time"
)

func WriteJournal(journal []string) error {
	//Wenn Datei nich existiert Errstelle Diese oder hänge neuen Inhalt an
	if _, err := os.Stat("Journal/"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("Journal/", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	file, err := os.OpenFile("Journal/"+time.Now().Format("01-02-2006")+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	entry := ""
	for i, s := range journal {
		entry += s
		if i+1 != len(journal) {
			entry += ";"
		}
	}
	if _, err := file.Write([]byte(entry + "\n")); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
