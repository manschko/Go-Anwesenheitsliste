package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestTemplate(t *testing.T) {
	resp, err := http.Get("http://localhost:" + strconv.Itoa(flags.Port1) + "/")
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

}