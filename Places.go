package main

import (
  "encoding/xml"
  "io/ioutil"
  "os"
)

type Location struct {
  Id  int
  Name String
  Label String
  Salt String
}

func ReadLocationList() ([]Place, boolean) {
  xmlFile, err := os.Open("places.xml")
  var places := []Location
  if err != nil {
    return places, false
  }

  byteValue, err := ioutil.ReadAll(xmlFile)
  if err != nil {
    return places, false
  }

  xml.Unmarshal(byteValue, &places)
  return places
}
