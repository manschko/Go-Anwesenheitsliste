package main

import (
  "encoding/xml"
  "io/ioutil"
  "os"
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

// This function read a XML file, which contains many locations
// Each location contains the name, the access token the current token and an old token
func ReadLocationList() ([]Location, bool) {
  // Open file
  xmlFile, err := os.Open("location.xml")
  if err != nil {
    return nil, false
  }

  // Reading file
  byteValue, err := ioutil.ReadAll(xmlFile)
  if err != nil {
    return nil, false
  }

  // Close file
  defer xmlFile.Close()

  // Load file content into Locations object
  var locations Locations
  xml.Unmarshal(byteValue, &locations)

  return locations.Locations, true
}

// This function writes the location configuration into a XML file
func  WriteLocationListToFile(locations []Location) bool {
  // Load a list of Location objects into a Locations object
  var location Locations
  location.Locations = locations

  // Convert object into string
  xmlString, err := xml.MarshalIndent(location, "", "  ")
  if err != nil {
    return false
  }

  // Write string into file
  os.WriteFile("location.xml", []byte(xmlString), 755)

  return true
}
