package main

import "testing"

func TestFlags(t *testing.T){
	//TODO os.agrs um flags zu setzen

	if flags.Port2 != 8080 {
		t.Errorf("flag for QR code Page expected 8080 got: %d", flags.Port2)
	}
	if flags.Port1 != 8000 {
		t.Errorf("flag for Login Page expected 8000 got: %d", flags.Port1)
	}
	if flags.TokenValidity != 5 {
		t.Errorf("flag for QR code Page expected 8080 got: %d", flags.TokenValidity)
	}

}