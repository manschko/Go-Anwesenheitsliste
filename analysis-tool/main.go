package main

import(
  "fmt"
  "bufio"
  "os"
  "strings"
  "io/ioutil"
  "regexp"
  "strconv"
)

// This function read the input of stdin
func ReadStdIn()(string, []string) {
  // Create new reader for stdin
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("-> ")

  // Read input
  text, _ := reader.ReadString('\n')

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
func PrintHelp() {
  fmt.Print("help\t\t\t\tZeigt diese Hilfe an\n")
  fmt.Print("\tselect-day 21-12-2012\t\tWählt ein Datum für weitere Befehle aus\n")
  fmt.Print("\tsearch-person max.mustermann\tSucht die Orte an dem sich eine Person am Tag aufgehalten hat\n")
  fmt.Print("\tlist-days\t\t\tZeigt die Tage an, an denen eine Anwesenheit protokolliert wurde\n")
  fmt.Print("\texport-list Ort\t\t\tExportiert die Anwesenheitsliste für einen Ort in eine CSV-Datei\n")
  fmt.Print("\texit\t\t\t\tBeendet dieses Programm")

  fmt.Print("\n")
}

// This function list all days, which have a log file
func ListDays() bool {

  // List all files in the journal directory
  files, err := ioutil.ReadDir("./Journal/")
  if (err != nil) {
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

    fmt.Print("\t" + strconv.Itoa(i + 1) + "\t| " + file.Name() + "\n")
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
  for _, row := range(rows) {
    fields := strings.Split(row, ";")
    if fields[0] != parameter[0] {
      continue
    }

    containPlace := false
    for _, place := range(places) {
      if fields[1] == place {
        containPlace = true
      }
    }

    if !containPlace {
      places = append(places, fields[1])
      fmt.Print("\t- " + fields[1] + "\n")
    }
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

  for _, row := range(rows) {
    fields := strings.Split(row, ";")
    if (len(fields) != 3) {
      continue
    }

    if fields[1] != parameter[0] {
      continue
    }

    containPerson := false
    for _, person := range(people) {
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

func main() {

  fmt.Print("\n### Go Anwesenheitsliste - Analyse Tool - begin ###\n\n")

  runCondition := true
  selectedDay := ""

  // Start program loop
  for runCondition {

    // Read stdin
    cmd, parameter := ReadStdIn()

    // Coordinate new command
    switch cmd {
      // Select a day
      case "select-day":
        if len(parameter) != 1 {
          PrintHelp()
          break
        }
        selectedDay = parameter[0]
        fmt.Print("Der " + parameter[0] + " wurde ausgewählt\n")
        break

      // List all days in the journal directory
      case "list-days":
        ListDays()
        break

      // Search a person in a log file
      case "search-person":
        if len(parameter) != 2 {
          PrintHelp()
          break
        }

        name := parameter[0] + " " + parameter[1]
        parameter := []string{name, selectedDay}
        SearchPerson(parameter)
        break

      // Export the attandance list into a file
      case "export-list":
        if len(parameter) == 0 {
          PrintHelp()
          break
        }

        place := ""
        for i, part := range(parameter) {
          place += part
          if i < len(parameter) - 1 {
            place += " "
          }
        }

        parameter := []string{place, selectedDay}
        ExportList(parameter)
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
