package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestQRSelectionPage(t *testing.T) {
	locations, _ := ReadLocationList()
	res := httptest.NewRecorder()
	//test if webpage is reachable
	req, err := http.NewRequest("GET", "https://"+flags.Url+":"+strconv.Itoa(flags.Port2)+"/", nil)
	selectionPageHandler(res, req)

	if err != nil {
		t.Fatal(err)

	}

	data := url.Values{}
	data.Set("location", locations[0].Name)
	req, err = http.NewRequest("POST", "https://"+flags.Url+":"+strconv.Itoa(flags.Port2)+"/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 || response.Request.URL.String() == "https://"+flags.Url+":"+strconv.Itoa(flags.Port2)+"/" {
		t.Errorf("Redirect to QR Page did not work")
	}

}
func TestQRPage(t *testing.T) {
	//check of redirect with wrong parameters
	locations, _ := ReadLocationList()
	req, err := http.NewRequest("GET", "https://"+flags.Url+":"+strconv.Itoa(flags.Port2)+"/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 || response.Request.URL.String() != "https://"+flags.Url+":"+strconv.Itoa(flags.Port2)+"/" {
		t.Errorf("Redirect with wrong Accescode did not work")
	}

	//check if QR code is getting generated
	os.Remove("PageTemplates/qr.png")

	req, err = http.NewRequest("GET", "https://"+flags.Url+":"+strconv.Itoa(flags.Port2)+"/"+locations[0].AccessToken, nil)

	if err != nil {
		t.Fatal(err)
	}

	response, err = client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println()

	_, err = os.Stat("PageTemplates/qr.png")
	errors.Is(err, os.ErrNotExist)
	if err != nil {
		t.Errorf("QR Code was not properly generated")
	}
}
