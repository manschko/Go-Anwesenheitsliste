package main

import (
	"flag"
	"net/http"
	"strconv"
	"testing"
)

//Überschreibe Main test Funktion um Funktionen vor den tests auszuführen
/*func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}*/




func TestWebServerAndFlags(t *testing.T) {
	ConfigFlag()
	//test for default Flags
	if flags.Port2 != 8080 {
		t.Errorf("flag for QR code Page expected 8080 got: %d", flags.Port2)
	}
	if flags.Port1 != 8000 {
		t.Errorf("flag for Login Page expected 8000 got: %d", flags.Port1)
	}
	if flags.TokenValidity != 5 {
		t.Errorf("flag for validation time of Tokens expected 5 got: %d", flags.TokenValidity)
	}

	//test for setting Falgs
	flag.Set("portLogin", "8001")
	flag.Set("portQR", "8002")
	flag.Set("valid", "80")

	if flags.Port2 != 8002 {
		t.Errorf("flag for QR code Page expected 8002 got: %d", flags.Port2)
	}
	if flags.Port1 != 8001 {
		t.Errorf("flag for Login Page expected 8001 got: %d", flags.Port1)
	}
	if flags.TokenValidity != 80 {
		t.Errorf("flag for validation time of Tokens expected 80 got: %d", flags.TokenValidity)
	}
	ConfigWebServer()
	resp, err := http.Get("http://localhost:" + strconv.Itoa(flags.Port2) + "/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Webserver mit dem port "+strconv.Itoa(flags.Port2)+" konnte nicht erreicht werden status code: %d", resp.StatusCode)
	}

	resp, err = http.Get("http://localhost:" + strconv.Itoa(flags.Port1) + "/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Webserver mit dem port "+strconv.Itoa(flags.Port1)+" konnte nicht erreicht werden status code: %d", resp.StatusCode)

	}


}