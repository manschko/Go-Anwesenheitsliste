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

// Test function CreateAndRunTimer
func TestCreateAndRunTimer(t *testing.T) {
	// Create parameter
	var mosbach Location
	mosbach.Id = 1
	mosbach.Name = "Mosbach"
	mosbach.AccessToken = "rg"
	mosbach.CurrentToken = "16192001216107769614"
	mosbach.OldToken = "12879960405315080572"

	var locationList []Location
	locationList = append(locationList, mosbach)

	result := CreateAndRunTimer(10, locationList)

	// Validate result
	if result == false {
		t.Error(result)
	}
}

// Test RunChangeTokenThread
func TestRunChangeTokenThread(t *testing.T) {
	// Run function
	result := RunChangeTokenThread()

	// Validate result
	if result == false {
		t.Error(result)
	}
}

// Test IsTokenValid
func TestIsTokenValid(t *testing.T) {
	// Create content for location list
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
	WriteLocationListToFile(locations)

	// Validate result
	result := IsTokenValid(mosbach.AccessToken, mosbach.OldToken)
	if result == false {
		t.Error(result)
	}
}
