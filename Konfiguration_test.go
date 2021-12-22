package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"os"
	"testing"
)

//Überschreibe Main test Funktion um Funktionen vor den tests auszuführen
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	ConfigFlag()
	ConfigWebServer()
	os.Rename("Journal", "JournalOld")
	os.Rename("JournalTest", "Journal")
}

func shutdown() {
	os.Rename("Journal", "JournalTest")
	os.Rename("JournalOld", "Journal")
}

func TestWebServerAndFlags(t *testing.T) {

	//test for default Flags
	if flags.Port2 != 8080 {
		t.Errorf("flag for QR code Page expected 8080 got: %d", flags.Port2)
	}
	if flags.Port1 != 8081 {
		t.Errorf("flag for Login Page expected 8000 got: %d", flags.Port1)
	}
	if flags.TokenValidity != 3600 {
		t.Errorf("flag for validation time of Tokens expected 3600 got: %d", flags.TokenValidity)
	}
	if flags.Url != "localhost" {
		t.Errorf("flag for URL expected localhost got: %s", flags.Url)
	}
}
