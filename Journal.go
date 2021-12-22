package main

import (
	"os"
	"time"
)

func WriteJournal(journal []string) error {
	//Wenn Datei nich existiert Errstelle Diese oder hänge neuen Inhalt an
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
