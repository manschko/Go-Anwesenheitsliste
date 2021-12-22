package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestQRSelectionPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(selectionPageHandler))
	locations, _ := ReadLocationList()
	//test if webpage is reachable
	_, err := http.Get(ts.URL)

	if err != nil {
		t.Fatal(err)

	}

	data := url.Values{}
	data.Set("location", locations[0].Name)
	res, err := http.PostForm(ts.URL, data)

	if err != nil {
		t.Fatal(err)
	}

	/*tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)

	res, err := http.Get(ts.Url)

	if err != nil {
		t.Fatal(err)
	}*/

	if res.StatusCode != 200 || res.Request.URL.String() == ts.URL {
		t.Errorf("Redirect to QR Page did not work")
	}

}

func TestQRPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(selectionPageHandler))
	//check of redirect with wrong parameters
	locations, _ := ReadLocationList()
	res, err := http.Get(ts.URL + "/test")
	if err != nil {
		t.Fatal(err)
	}
	/*tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	response, err := client.Do(req)
	*/
	if err != nil {
		t.Fatal(err)
	}
	//test redirect
	if res.StatusCode != 200 || res.Request.URL.String() != ts.URL+"/" {
		t.Errorf("Redirect with wrong Accescode did not work")
	}

	//check if QR code is getting generated
	os.Remove("PageTemplates/qr.png")

	_, err = http.Get(ts.URL + "/" + locations[0].AccessToken)

	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat("PageTemplates/qr.png")
	errors.Is(err, os.ErrNotExist)
	if err != nil {
		t.Errorf("QR Code was not properly generated")
	}
}
