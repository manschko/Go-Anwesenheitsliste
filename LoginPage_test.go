package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestToken(t *testing.T) {
	locations, _ := ReadLocationList()
	res := httptest.NewRecorder()
	//test für error handling bei keinem übergebenen Token
	req, err := http.NewRequest("GET", "https://"+flags.Url+":"+strconv.Itoa(flags.Port1)+"/", nil)
	LoginPageHandler(res, req)

	if err != nil {
		t.Fatal(err)
	}
	if !templateDataLogin.Failed {
		t.Errorf("Token ist vorhanden wurde aber nicht übergeben")
	}
	//test ob Tokens übetragen und erkannt wurden
	req, err = http.NewRequest("GET", "https://"+flags.Url+""+strconv.Itoa(flags.Port1)+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, nil)
	LoginPageHandler(res, req)
	if err != nil {
		t.Fatal(err)
	}
	if templateDataLogin.Failed {
		t.Errorf("Es konnten keine Token gefunden werden")
	}

}

/*
LoginPage_test.go:50: Get "https://localhost:8081?location=rg&access=3890372420546292004": dial tcp 127.0.0.1:8081: connect: connection refused
Lokal funktioniert der Test, siehe PassedTests/TestForm.JPG bekomme aber auf der CLI connection refused
func TestForm(t *testing.T) {
	locations, _ := ReadLocationList()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get("https://" + flags.Url + ":" + strconv.Itoa(flags.Port1) + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.PostForm("https://"+flags.Url+":"+strconv.Itoa(flags.Port1)+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, url.Values{"adresse": {"test"}, "name": {"name"}, "submit": {"true"}})
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
	_, err = client.Get("https://" + flags.Url + ":" + strconv.Itoa(flags.Port1) + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	if templateDataLogin.Adresse != "test" || templateDataLogin.Name != "name" {
		t.Errorf("Daten konten nicht von der Map wieder hergestellt werden")
	}
	//check ob Mapeintrag gelöscht wird wenn Token ungültig
	RunChangeTokenThread()
	RunChangeTokenThread()
	_, err = client.Get("https://" + flags.Url + ":" + strconv.Itoa(flags.Port1) + "?location=" + locations[0].AccessToken + "&access=" + locations[0].CurrentToken)
	if err != nil {
		t.Fatal(err)
	}
	if templateDataLogin.Adresse != "" || templateDataLogin.Name != "" {
		t.Errorf("Daten wurden mit Ungültikem Token wiederhergestellt")
	}
	//check ob Mapeintrag gelöscht wrid wenn Abgemeldet
	locations, _ = ReadLocationList()
	_, err = client.PostForm("https://"+flags.Url+":"+strconv.Itoa(flags.Port1)+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, url.Values{"adresse": {"test"}, "name": {"name"}, "submit": {"true"}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.PostForm("https://"+flags.Url+":"+strconv.Itoa(flags.Port1)+"?location="+locations[0].AccessToken+"&access="+locations[0].CurrentToken, url.Values{"adresse": {"test"}, "name": {"name"}, "submit": {"false"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(dataMap) != 0 {
		t.Errorf("Eintrag wurde nicht aus der Map gelöscht nach dem Abmelden")
	}
}*/
