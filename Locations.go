package main

import (
  "encoding/xml"
  "io/ioutil"
  "os"
  "fmt"
)

type Locations struct {
  Locations []Location `xml:"location"`
}

type Location struct {
  Id int `xml:"id"`
  Name string `xml:"name"`
  AccessToken string `xml:"accesstoken"`
  CurrentToken string
  OldToken string
}

// Mit dieser Funktion bekommt man den Inhalt von location.xml als Objekt
// und kann den Zugnag Token und den aktuellen und vorherigen Zeit Token auslesen
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

// Diese Funktion schreibt die geänderten Werte zurück in die Datei location.xml
func  WriteLocationListToFile(locations []Location) {
  var location Locations
  location.Locations = locations

  xmlString, err := xml.MarshalIndent(location, "", "  ")
  if err != nil {
    return
  }

  fmt.Print(string(xmlString), "\n\n")
  os.WriteFile("location.xml", []byte(xmlString), 755)
}

func printLocation(location Location) {
  fmt.Printf("\nId:%v\nName:%v\nCurrent token:%v\nOld token:%v", location.Id, location.Name, location.CurrentToken, location.OldToken)
}
