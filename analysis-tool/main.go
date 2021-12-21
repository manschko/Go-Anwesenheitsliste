package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// This function read the input of stdin
func ReadStdIn(text string) (string, []string) {

	// Trim LF
	text = strings.Trim(text, "\n")
	textSplit := strings.Split(text, " ")

	// Split input into command and parameter
	command := textSplit[0]
	parameter := textSplit[1:]

	fmt.Print("out:\t")

	return command, parameter
}

// This function read the complete content of a journal
func GetFileContent(day string) string {
	// Read a journal file
	byteContent, err := os.ReadFile("Journal/" + day)

	if err != nil {
		return "error"
	}

	// Convert byte content into string
	content := string(byteContent)

	return content
}

// This function print a help into a terminal

func PrintHelp() bool {
	fmt.Print("help\t\t\t\tZeigt diese Hilfe an\n")
	fmt.Print("\tselect-day 21-12-2012\t\tWählt ein Datum für weitere Befehle aus\n")
	fmt.Print("\tsearch-person max mustermann\tSucht die Orte an dem sich eine Person am Tag aufgehalten hat\n")
  	fmt.Print("\tcontact-list max mustermann\tZeigt die Kontakt Zeit die diese Person mit anderne hatte\n")
	fmt.Print("\tlist-days\t\t\tZeigt die Tage an, an denen eine Anwesenheit protokolliert wurde\n")
	fmt.Print("\texport-list Ort\t\t\tExportiert die Anwesenheitsliste für einen Ort in eine CSV-Datei\n")
	fmt.Print("\texit\t\t\t\tBeendet dieses Programm")

	fmt.Print("\n")

	return true
}

// This function list all days, which have a log file
func ListDays() bool {

	// List all files in the journal directory
	files, err := ioutil.ReadDir("./Journal/")
	if err != nil {
		return false
	}

	// Use a regex expression to get
	regexp := regexp.MustCompile("[0-9][0-9]-[0-9][0-9]-[0-9][0-9][0-9][0-9]")
	fmt.Print("Index\t| Tag\n")

	// Print file name
	for i, file := range files {
		result := regexp.FindString(file.Name())

		if result == "" {
			continue
		}

		fmt.Print("\t" + strconv.Itoa(i+1) + "\t| " + file.Name() + "\n")
	}

	fmt.Print("\n")

	return true
}

// This function find login data from a person in a specific day
func SearchPerson(parameter []string) bool {
	// Get content of a journal
	content := GetFileContent(parameter[1])
	if content == "no result" {
		return false
	}

	rows := strings.Split(content, "\n")
	var places []string

	// Print data of a person
	fmt.Print("Gesuchter Name: " + parameter[0] + "\n")
	fmt.Print("\tOrte:\n")
	for _, row := range rows {
		fields := strings.Split(row, ";")
		if fields[0] != parameter[0] {
			continue
		}

		containPlace := false
		for _, place := range places {
			if fields[1] == place {
				containPlace = true
			}
		}

		if !containPlace {
			places = append(places, fields[1])
			fmt.Print("\t- " + fields[1] + "\n")
		}
	}

	if len(places) == 0 {
		return false
	}
	return true
}

// This function export a attandance list of location into a CSV file
func ExportList(parameter []string) bool {

	// Get content of a file
	content := GetFileContent(parameter[1])
	rows := strings.Split(content, "\n")
	var people []string
	index := 1

	// Create file
	file, err := os.Create(parameter[1] + "-" + parameter[0] + "-export.csv")
	if err != nil {
		return false
	}

	// Create new file writer
	writer := bufio.NewWriter(file)
	fmt.Print("Index\t| Name\n")

	for _, row := range rows {
		fields := strings.Split(row, ";")
		if len(fields) != 3 {
			continue
		}

		if fields[1] != parameter[0] {
			continue
		}

		containPerson := false
		for _, person := range people {
			if fields[0] == person {
				containPerson = true
			}
		}

		if !containPerson {
			people = append(people, fields[0])
			fmt.Print("\t" + strconv.Itoa(index) + "\t| " + fields[0] + "\n")

			// Write data set into file
			writer.WriteString(strconv.Itoa(index) + ";" + fields[0] + "\n")
		}
	}

	writer.Flush()

	file.Sync()
	file.Close()

	fmt.Print(
		"\n\tDie Liste wurde in ./" +
			parameter[1] +
			"-" +
			parameter[0] +
			"-export.csv" +
			" exportiert\n")

	return true
}

func SearchContact(name string) {
	isCheckedIn := false
	currentVisitors := make(map[string]map[string]time.Time)
	contact := make(map[string]time.Duration)
	suspect := struct {
		n        string
		location string
	}{name, ""}
	files, err := ioutil.ReadDir("Journal")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		content := GetFileContent(file.Name())
		rows := strings.Split(content, "\n")

		for _, row := range rows {
			c := strings.Split(row, ";")

			if c[2] == suspect.n {
				if c[3] == "Anmeldung" {
					suspect.location = c[0]
					isCheckedIn = true
				} else {
					isCheckedIn = false
				}
			}

			if _, ok := currentVisitors[c[0]]; !ok {
				currentVisitors[c[0]] = make(map[string]time.Time)
			}
			if _, ok := currentVisitors[c[0]][c[2]]; !ok {
				
				t, _ := time.Parse("01-02-200615:04", strings.Split(file.Name(), ".")[0]+c[4])
				currentVisitors[c[0]][c[2]] = t
			} else if c[3] == "Abmeldung" {
				if suspect.n != c[2] && isCheckedIn && suspect.location == c[0] {
					//check if person checked in after suspect
					if currentVisitors[c[0]][suspect.n].After(currentVisitors[c[0]][c[2]]) {
						//person checkin time - checkoutime
						if _, ok := contact[c[2]]; !ok {
							t, _ := time.Parse("01-02-200615:04", strings.Split(file.Name(), ".")[0]+c[4])
							contact[c[2]] = t.Sub(currentVisitors[c[0]][c[2]])
						} else {
							t, _ := time.Parse("01-02-200615:04", strings.Split(file.Name(), ".")[0]+c[4])
							contact[c[2]] += t.Sub(currentVisitors[c[0]][c[2]])
						}
						//if Person checked in before suspect
						//Person checkouttime - suspect checkintime
					} else {
						t, _ := time.Parse("01-02-200615:04", strings.Split(file.Name(), ".")[0]+c[4])
						contact[c[2]] = currentVisitors[c[0]][name].Sub(t)
					}
				} else if suspect.n == c[2] {
					for key, value := range currentVisitors[c[0]] {
						if key != c[2] {
							//check if visitor checked in after suspect
							if currentVisitors[c[0]][key].After(currentVisitors[c[0]][c[2]]) {
								t, _ := time.Parse("01-02-200615:04", strings.Split(file.Name(), ".")[0]+c[4])
								if _, ok := contact[key]; !ok {
									contact[key] = t.Sub(value)
								} else {
									contact[key] += t.Sub(value)
								}
							} else {
								t, _ := time.Parse("01-02-200615:04", strings.Split(file.Name(), ".")[0]+c[4])
								if _, ok := contact[key]; !ok {
									contact[key] = t.Sub(currentVisitors[c[0]][c[2]])
								} else {
									contact[key] += t.Sub(currentVisitors[c[0]][c[2]])
								}
							}
						}
					}
				}
				delete(currentVisitors[c[0]], c[2])
			}
		}
		fmt.Println("\nKontakt dauer zu verdächtigen " + suspect.n)
		for key, value := range contact {
			fmt.Println(key + ": " + value.String())
		}
	}
}
func ExecSelectDay(selectedDay string, parameter []string) bool {
	if len(parameter) != 1 {
		PrintHelp()
		return false
	}
	selectedDay = parameter[0]
	fmt.Print("Der " + parameter[0] + " wurde ausgewählt\n")
	return true
}

func ExecSearchPerson(selectedDay string, parameter []string) bool {
	if len(parameter) != 2 {
		PrintHelp()
		return false
	}

	name := parameter[0] + " " + parameter[1]
	parameter = []string{name, selectedDay}
	SearchPerson(parameter)

	return true
}

func ExecExportList(selectedDay string, parameter []string) bool {
	if len(parameter) == 0 {
		PrintHelp()
		return false
	}

	place := ""
	for i, part := range parameter {
		place += part
		if i < len(parameter)-1 {
			place += " "
		}
	}

	parameter = []string{place, selectedDay}
	ExportList(parameter)
	return true
}

func main() {

	fmt.Print("\n### Go Anwesenheitsliste - Analyse Tool - begin ###\n\n")

	runCondition := true
	selectedDay := ""

	// Start program loop
	for runCondition {

		// Read stdin

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")

		// Read input
		text, _ := reader.ReadString('\n')
		cmd, parameter := ReadStdIn(text)

		// Coordinate new command
		switch cmd {
		// Select a day
		case "select-day":
			ExecSelectDay(selectedDay, parameter)
			break

		// List all days in the journal directory
		case "list-days":
			ListDays()
			break

		// Search a person in a log file
		case "search-person":
			ExecSearchPerson(selectedDay, parameter)
			break

		// Export the attandance list into a file
		case "export-list":

      //TODO implement
			if len(parameter) == 0 {
				PrintHelp()
				break
			}

			place := ""
			for i, part := range parameter {
				place += part
				if i < len(parameter)-1 {
					place += " "
				}
			}

			parameter := []string{place, selectedDay}
			ExportList(parameter)
			break

		//Export contact list
		case "contact-list":
			if len(parameter) != 2 {
				PrintHelp()
				break
			}
    
			name := parameter[0] + " " + parameter[1]
			SearchContact(name)
			break
      
		// Stop the analyse program
		case "exit":
			fmt.Print("\tDas Programm wird beendet\n")
			runCondition = false

		default:
			PrintHelp()
		}

		fmt.Print("\n")
	}

	fmt.Print("### Go Anwesenheitsliste - Analyse Tool - end ###\n\n")
}


