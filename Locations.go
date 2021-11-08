package main

import (
  "encoding/xml"
  "io/ioutil"
  "os"
)

type Locations struct {
  Test string `xml:"test"`
  Locations []Location `xml:"location"`
}

type Location struct {
  Id int `xml:"id"`
  Name string `xml:"name"`
}

func ReadLocationList() ([]Location, bool) {
  xmlFile, err := os.Open("location.xml")
  if err != nil {
    return nil, false
  }

  byteValue, err := ioutil.ReadAll(xmlFile)
  if err != nil {
    return nil, false
  }

  defer xmlFile.Close()

  var locations Locations
  xml.Unmarshal(byteValue, &locations)

  return locations.Locations, true
}
