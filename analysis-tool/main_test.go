package main

import (
	"os"
	"testing"
)

func TestReadStdIn(t *testing.T) {
	command, parameter := ReadStdIn("help")
	if command != "help" || len(parameter) != 0 {
		t.Error(command, parameter)
	}
}

// Test GetFileContent
func TestGetFileContent(t *testing.T) {
	// Change working directory
	os.Chdir("../")

	// Execute function
	content := GetFileContent("12-07-2021.txt")

	// Validate content
	if content == "error" {
		t.Error(content)
	}

	if len(content) == 0 {
		t.Error(content)
	}
}

func TestPrintHelp(t *testing.T) {
	if PrintHelp() == false {
		t.Error("Printing of help text fails")
	}
}

func TestListDays(t *testing.T) {
	result := ListDays()

	if result == false {
		t.Error("Error in getting list of available days")
	}
}

func TestSearchPerson(t *testing.T) {
	var parameter []string
	parameter = append(parameter, "name")
	parameter = append(parameter, "12-07.2021")
	result := SearchPerson(parameter)

	if result == true {
		t.Error("User does not login in or logout on this day")
	}
}

func TestExportList(t *testing.T) {
	var parameter []string
	parameter = append(parameter, "name")
	parameter = append(parameter, "12-07.2021")
	result := ExportList(parameter)

	if result == false {
		t.Error("Extraction of person fails")
	}

	byteContent, error := os.ReadFile("12-07.2021-name-export.csv")

	if error != nil {
		t.Error("Reading of export fails")
	}

	if len(byteContent) != 0 {
		t.Error("Find content but no user login or logout on this day")
	}
}

func TestExecSelectDay(t *testing.T) {
	var parameter []string
	parameter = append(parameter, "12-07-2021")
	var selecteDay string
	result := ExecSelectDay(selecteDay, parameter)

	if result == false {
		t.Error("Parameter contains no day")
	}
}

func TestExecSearchPerson(t *testing.T) {
	selectedDay := "12-07-2021"
	var parameter []string
	parameter = append(parameter, "name")
	parameter = append(parameter, "12-07.2021")
	result := ExecSearchPerson(selectedDay, parameter)
	if result == false {
		t.Error("Not defined number of parameter")
	}
}

func TestExecExportList(t *testing.T) {
	selectedDay := "12-07-2021"
	var parameter []string
	parameter = append(parameter, selectedDay)
	parameter = append(parameter, "Mosbach")

	if !ExecExportList(selectedDay, parameter) {
		t.Error("Execution fails")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
