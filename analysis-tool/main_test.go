package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	os.Chdir("..")
	// Execute function
	content := GetFileContent("12-20-2021")

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
	if SearchPerson([]string{"", ""}) {
		t.Error("Unexpected parameters")
	}
	parameter = append(parameter, "test2 test")
	parameter = append(parameter, "12-19-2021")

	if !SearchPerson(parameter) {
		t.Error("Did not find expected user")
	}

	parameter[0] = "test4 test"
	if SearchPerson(parameter) {
		t.Error("User does not login in or logout on this day")
	}
}

func TestExportList(t *testing.T) {
	var parameter []string
	parameter = append(parameter, "Bad Mergendheim")
	parameter = append(parameter, "12-19-2021")
	result := ExportList(parameter)

	if result == false {
		t.Error("Extraction of person fails")
	}
	byteContent, error := os.ReadFile("12-19-2021-Bad Mergendheim-export.csv")

	fmt.Println(os.Getwd())
	if error != nil {
		t.Error(error)
	}

	if len(byteContent) == 0 {
		t.Error("Find content but no user login or logout on this day")
	}
}

func TestExecSelectDay(t *testing.T) {
	var parameter []string
	parameter = append(parameter, "12-19-2021")
	var selecteDay string

	if !ExecSelectDay(selecteDay, parameter) {
		t.Error("Parameter contains no day")
	}

	if ExecSelectDay(selecteDay, []string{}) {
		t.Error("Invalid Parameter")
	}
}

func TestExecSearchPerson(t *testing.T) {
	selectedDay := "12-07-2021"
	var parameter []string
	parameter = append(parameter, "name")
	parameter = append(parameter, "12-07.2021")
	if !ExecSearchPerson(selectedDay, parameter) {
		t.Error("Not defined number of parameter")
	}

	if ExecSearchPerson(selectedDay, []string{""}) {
		t.Error("Invalid Parameter")
	}
}

func TestExecExportList(t *testing.T) {
	selectedDay := "12-07-2021"
	var parameter []string
	parameter = append(parameter, selectedDay)
	parameter = append(parameter, "Mosbach")

	if !ExecExportList(selectedDay, parameter) {
		t.Error("Invalid Parameter")
	}

	if ExecExportList(selectedDay, []string{}) {
		t.Error("Invalid Parameter")
	}

	if ExecExportList(selectedDay, []string{"", "", ""}) {
		t.Error("Invalid Parameter")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
func TestContactList(t *testing.T) {
	output := SearchContact("test test")
	rows := strings.Split(output, "\n")
	rows = rows[2:]

	expected := []string{"test3 test: 3h56m0s", "test2 test: 3h58m0s", "test4 test: 4h0m0s", "test5 test: 4h0m0s", "test7 test: 3h54m0s", "test6 test: 3h54m0s", "test9 test: 2h2m0s", "test8 test: 2h1m0s"}
	//sort both slices to get two equal slices
	sort.Strings(expected)
	sort.Strings(rows)
	//check if both slices equal length
	if len(rows) != len(expected) {
		t.Error("output length: " + strconv.Itoa(len(rows)) + " is not the same as expected: " + strconv.Itoa(len(expected)))
	}

	//check for differences in slice
	for i, row := range rows {
		if expected[i] != row {
			t.Error("Missmatch fount between output: " + row + " and expected value: " + expected[i])
		}
	}
}
