package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWriteJournal(t *testing.T) {
	WriteJournal("test", "testort", true)
	file, err := os.OpenFile("Journal/" + time.Now().Format("01-02-2006") + ".txt", os.O_RDONLY, 0660)

	if err != nil {
		t.Errorf("error beim öffnen der Datei: %d", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lastLine string

	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	s := strings.Split(lastLine, ";")

	if s[0] != "test" {
		t.Errorf("name in Journal expected test got: %s", s[0])
	}

	if s[1] != "testort" {
		t.Errorf("place in Journal expected testort got: %s", s[1])
	}

	if s[2] != "true" {
		t.Errorf("Login type in Journal expected true got: %s", s[3])
	}

}