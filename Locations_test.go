package main

import (
  "testing"
)

func TestWriteLocationListToFile(t *testing.T) {
  var mosbach Location
  mosbach.Id = 1
  mosbach.Name = "Mosbach"
  mosbach.AccessToken = "rg"
  mosbach.CurrentToken = "3890372420546292004"
  mosbach.OldToken = "25669930129342924"

  var badMergendheim Location
  badMergendheim.Id = 2
  badMergendheim.Name = "Bad Mergendheim"
  badMergendheim.AccessToken = "erg"
  badMergendheim.CurrentToken = "11505544865499498832"
  badMergendheim.OldToken = "17337374235514595919"

  var locations []Location
  locations = append(locations, mosbach)
  locations = append(locations, badMergendheim)

  result := WriteLocationListToFile(locations)

  if result == false {
    t.Error(result)
  }
}

func TestReadLocationList(t *testing.T) {
  locations, result := ReadLocationList()

  if locations == nil {
    t.Error(locations)
  }

  if result == false {
    t.Error(result)
  }
}
