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

func ReadStdIn()(string, []string) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("-> ")
  text, _ := reader.ReadString('\n')

  text = strings.Trim(text, "\n")
  textSplit := strings.Split(text, " ")
  command := textSplit[0]
  parameter := textSplit[1:]

  fmt.Print("out:\t")

  return command, parameter
}

func GetFileContent(day string) string {
  byteContent, err := os.ReadFile("Journal/" + day)

  if err != nil {
    return "error"
  }

  content := string(byteContent)

  return content
}

func PrintHelp() {
  fmt.Print("help\t\t\t\tZeigt diese Hilfe an\n")
  fmt.Print("\tselect-day 21-12-2012\t\tWählt ein Datum für weitere Befehle aus\n")
  fmt.Print("\tsearch-person max.mustermann\tSucht die Orte an dem sich eine Person am Tag aufgehalten hat\n")
  fmt.Print("\tlist-days\t\t\tZeigt die Tage an, an denen eine Anwesenheit protokolliert wurde\n")
  fmt.Print("\texport-list Ort\t\t\tExportiert die Anwesenheitsliste für einen Ort in eine CSV-Datei\n")
  fmt.Print("\texit\t\t\t\tBeendet dieses Programm")

  fmt.Print("\n")
}

func ListDays() {

  files, err := ioutil.ReadDir("Journal/")

  if (err != nil) {
    return
  }

  regexp := regexp.MustCompile("[0-9][0-9]-[0-9][0-9]-[0-9][0-9][0-9][0-9]")
  fmt.Print("Index\t| Tag\n")

  for i, file := range files {
    result := regexp.FindString(file.Name())

    if result == "" {
      fmt.Print("we")
      continue
    }

    fmt.Print("\t" + strconv.Itoa(i + 1) + "\t| " + file.Name() + "\n")
  }

  fmt.Print("\n")
}

func SearchPerson(parameter []string) {
  content := GetFileContent(parameter[1])
  rows := strings.Split(content, "\n")
  var places []string

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
}

func ExportList(parameter []string) {
  content := GetFileContent(parameter[1])
  rows := strings.Split(content, "\n")
  var people []string
  index := 1

  file, err := os.Create(parameter[1] + "-" + parameter[0] + "-export.csv")
  if err != nil {
    return
  }

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
}

func main() {

  fmt.Print("\n### Go Anwesenheitsliste - Analyse Tool - begin ###\n\n")

  runCondition := true
  selectedDay := ""

  for runCondition {

    cmd, parameter := ReadStdIn()

    switch cmd {

      case "select-day":
        if len(parameter) != 1 {
          PrintHelp()
          break
        }
        selectedDay = parameter[0]
        fmt.Print("Der " + parameter[0] + " wurde ausgewählt\n")
        break

      case "list-days":
        ListDays()
        break

      case "search-person":
        if len(parameter) != 2 {
          PrintHelp()
          break
        }

        name := parameter[0] + " " + parameter[1]
        parameter := []string{name, selectedDay}
        SearchPerson(parameter)
        break

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
