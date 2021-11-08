package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestTemplate(t *testing.T) {
	resp, err := http.Get("https://localhost:" + strconv.Itoa(flags.Port1) + "/")
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	//TODO compare body with template
	fmt.Println(string(body))
}
func TestForm(t *testing.T) {
	_, err := http.PostForm("https://localhost:" + strconv.Itoa(flags.Port1) + "/", url.Values{"adresse": {"test"}, "name": {"name"}})
	if err != nil {
		t.Fatal(err)
	}
	if loginData.Name != "name" {
		t.Errorf("Fehler bei dem übertragen von Formdaten erwartet \"name\" für den namen, übertragen: %s", loginData.Name)
	}
	if loginData.Adresse != "test" {
		t.Errorf("Fehler bei dem übertragen von Formdaten erwartet \"test\" für den adresse, übertragen: %s", loginData.Adresse)
	}
}