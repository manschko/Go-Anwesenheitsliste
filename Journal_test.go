package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWriteJournal(t *testing.T) {
	entry := []string{"testort", "testadresse", "test name", "Anmeldung","01:20"}
	os.Rename("Journal", "JournalTemp")
	err := WriteJournal(entry)
	//test für error handling in journal funktion
	if err != nil {
		t.Error(err)
	}
	os.Rename("JournalTemp","Journal")
	err = WriteJournal(entry)
	file, err := os.OpenFile("Journal/" + time.Now().Format("01-02-2006") + ".txt", os.O_RDONLY, 0660)
	//check ob auf schreiben der Datei geklappt hat
	if err != nil {
		t.Errorf("error beim öffnen der Datei: %d", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lastLine string
	//lese den Letzten Eintrag ein
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	s := strings.Split(lastLine, ";")
	//test ob der Eintrag mit dem aus dem Journal übereinstimmt
	for i, e := range entry {
		if s[i] != e {
			t.Errorf("error in Journal expected %s got: %s",e, s[0])
		}
	}
}