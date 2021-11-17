package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestToken(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "https://localhost:" + strconv.Itoa(flags.Port1) + "/", nil)
	LoginPageHandler(res, req)

	if err != nil{
		t.Fatal(err)
	}
	if !templateDataLogin.Failed {
		t.Errorf("Token ist vorhanden wurde aber nicht übergeben")
	}
	//todo key anpassen
	req, err = http.NewRequest("GET","https://localhost:" + strconv.Itoa(flags.Port1) + "/?access=test", nil)
	LoginPageHandler(res, req)
	if err != nil{
		t.Fatal(err)
	}
	/*if templateDataLogin. {
		t.Errorf("Falscher Token wurde erfasst. Erwatet: test, bekommen: %s", templateDataLogin.Key )
	}*/
}
func TestForm(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get("https://localhost:" + strconv.Itoa(flags.Port1) + "/")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.PostForm("https://localhost:" + strconv.Itoa(flags.Port1) + "/", url.Values{"adresse": {"test"}, "name": {"name"}})
	if err != nil {
		t.Fatal(err)
	}
	if loginData.Name != "name" {
		t.Errorf("Fehler bei dem übertragen von Formdaten erwartet \"name\" für den Namen, übertragen: %s", loginData.Name)
	}
	if loginData.Adresse != "test" {
		t.Errorf("Fehler bei dem übertragen von Formdaten erwartet \"test\" für die Adresse, übertragen: %s", loginData.Adresse)
	}
}


/*TODO
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


 */