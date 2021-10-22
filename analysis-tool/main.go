package main

import(
  "fmt"
  "bufio"
  "os"
  "strings"
  "io/ioutil"
  "regexp"
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
  fmt.Print("\thelp\t\t\t\tZeigt diese Hilfe an\n")
  fmt.Print("\t\tselect-day 21-12-2012\t\tWählt ein Datum für weitere Befehle aus\n")
  fmt.Print("\t\tsearch-person max.mustermann\tSucht die Orte an dem sich eine Person am Tag aufgehalten hat\n")
  fmt.Print("\t\tlist-days\t\t\tZeigt die Tage an, an denen eine Anwesenheit protokolliert wurde\n")
  fmt.Print("\t\texport-list Ort\t\t\tExportiert die Anwesenheitsliste für einen Ort in eine CSV-Datei\n")
  fmt.Print("\t\texit\t\t\t\tBeendet dieses Programm")

  fmt.Print("\n")
}

func ListDays() {

  files, err := ioutil.ReadDir("Journal/")

  if (err != nil) {
    return
  }

  regexp := regexp.MustCompile("[0-9][0-9]-[0-9][0-9]-[0-9][0-9][0-9][0-9]")
  fmt.Print("\tIndex\t| Tag\n")

  for i, file := range files {
    result := regexp.FindString(file.Name())

    if result == "" {
      fmt.Print("we")
      continue
    }

    fmt.Print("\t\t")
    fmt.Print(i + 1)
    fmt.Print("\t| " + file.Name() + "\n")
  }

  fmt.Print("\n")
}

func SearchPerson(parameter []string) {
  content := GetFileContent(parameter[1])
  rows := strings.Split(content, "\n")
  var places []string

  fmt.Print("\tGesuchter Name: " + parameter[0] + "\n")
  fmt.Print("\t\tOrte:\n")
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
      fmt.Print("\t\t- " + fields[1] + "\n")
    }
  }
}

func ExportList(parameter []string) {
  //content := GetFileContent(parameter[1])
  //rows := strings.Split(content, "\n")
  //var places []string
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
        fmt.Print("\tDer " + parameter[0] + " wurde ausgewählt\n")
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
        if len(parameter) != 1 {
          PrintHelp()
          break
        }

        parameter = append(parameter, selectedDay)
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
