package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(LoginPageHandler))
	locations, _ := ReadLocationList()
	//test für error handling bei keinem übergebenen Token
	_, err := http.Get(ts.URL)

	if err != nil {
		t.Fatal(err)
	}
	if !templateDataLogin.Failed {
		t.Errorf("Token ist vorhanden wurde aber nicht übergeben")
	}
	//test ob Tokens übetragen und erkannt wurden
	_, err = http.Get(ts.URL + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	if templateDataLogin.Failed {
		t.Errorf("Es konnten keine Token gefunden werden")
	}

}

func TestForm(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(LoginPageHandler))
	locations, _ := ReadLocationList()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get(ts.URL + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.PostForm(ts.URL+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, url.Values{"adresse": {"test"}, "name": {"name"}, "submit": {"true"}})
	if err != nil {
		t.Fatal(err)
	}
	if loginData.Name != "name" {
		t.Errorf("Fehler bei dem übertragen von Formdaten erwartet \"name\" für den Namen, übertragen: %s", loginData.Name)
	}
	if loginData.Adresse != "test" {
		t.Errorf("Fehler bei dem übertragen von Formdaten erwartet \"test\" für die Adresse, übertragen: %s", loginData.Adresse)
	}
	//check ob from Daten gespeichert wurden
	_, err = client.Get(ts.URL + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	if templateDataLogin.Adresse != "test" || templateDataLogin.Name != "name" {
		t.Errorf("Daten konten nicht von der Map wieder hergestellt werden")
	}
	//check ob Mapeintrag gelöscht wird wenn Token ungültig
	RunChangeTokenThread()
	RunChangeTokenThread()
	_, err = client.Get(ts.URL + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	if templateDataLogin.Adresse != "" || templateDataLogin.Name != "" {
		t.Errorf("Daten wurden mit Ungültikem Token wiederhergestellt")
	}
	//check ob Mapeintrag gelöscht wrid wenn Abgemeldet
	locations, _ = ReadLocationList()
	_, err = client.PostForm(ts.URL+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, url.Values{"adresse": {"test"}, "name": {"name"}, "submit": {"true"}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.PostForm(ts.URL+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, url.Values{"adresse": {"test"}, "name": {"name"}, "submit": {"false"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(dataMap) != 0 {
		t.Errorf("Eintrag wurde nicht aus der Map gelöscht nach dem Abmelden")
	}
}
