package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"testing"
)

// Test WriteLocationListToFile
func TestWriteLocationListToFile(t *testing.T) {
	// Create parameter of function WriteLocationListToFile
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

	// Write location list to file
	result := WriteLocationListToFile(locations)

	// Validate result
	if result == false {
		t.Error(result)
	}
}

// Test ReadLocationList
func TestReadLocationList(t *testing.T) {
	// Read locations
	locations, result := ReadLocationList()

	// Validate locations
	if locations == nil {
		t.Error(locations)
	}

	// Validate result
	if result == false {
		t.Error(result)
	}
}
