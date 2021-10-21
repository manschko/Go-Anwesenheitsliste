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

func GetFileReader(day string) []byte {
  reader, err := os.ReadFile("Journal/" + day)

  if err != nil {
    return nil
  }

  return reader
}

func ListDays() {

  files, err := ioutil.ReadDir("Journal/")

  if (err != nil) {
    return
  }

  regexp := regexp.MustCompile("[0-9]*-[0-9]*-[0-9]*")

  for _, file := range files {
    result := regexp.FindString(file.Name())

    if result == "" {
      fmt.Print("we")
      continue
    }

    fmt.Print("\t" + file.Name() + "\n")
  }

  fmt.Print("\n")
}

func SearchPerson(parameter []string) {

}

func ExportList(parameter []string) {

}

func main() {

  fmt.Print("\n### Go Anwesenheitsliste - Analyse Tool - begin ###\n\n")

  runCondition := true
  // selectedDay := ""

  for runCondition {

    cmd, parameter := ReadStdIn()

    switch cmd {

      case "select-day":

        fmt.Print("\tDer " + parameter[0] + " wurde ausgewählt\n")
        break

      case "list-days":
        ListDays()
        break

      case "search-person":
        SearchPerson(parameter)
        break

      case "export-list":
        ExportList(parameter)
        break

      case "exit":
        fmt.Print("\tDas Programm wird beendet\n")
        runCondition = false

      default:
        fmt.Print("\thelp\t\t\t\tZeigt diese Hilfe an\n")
        fmt.Print("\t\tselect-day 21-12-2012\t\tWählt ein Datum für weitere Befehle aus\n")
        fmt.Print("\t\tsearch-person max.mustermann\tSucht die Orte an dem sich eine Person am Tag aufgehalten hat\n")
        fmt.Print("\t\tlist-days\t\t\tZeigt die Tage an, an denen eine Anwesenheit protokolliert wurde\n")
        fmt.Print("\t\texport-list Ort\t\t\tExportiert die Anwesenheitsliste für einen Ort in eine CSV-Datei\n")
        fmt.Print("\t\texit\t\t\t\tBeendet dieses Programm")

        fmt.Print("\n")
    }

    fmt.Print("\n")
  }

  //selectedDay = "wef"

  fmt.Print("### Go Anwesenheitsliste - Analyse Tool - end ###\n\n")
}
