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

  fmt.Print("out:\n")

  return command, parameter
}

func ListDays() {

  files, err := ioutil.ReadDir("./Journal/")
  if (err == nil) {

    return
  }

  for _, file := range files {
    result, err := regexp.Match(".[0-9]*-.[0-9]*-.[0-9]*", []byte(file.Name()))

    if err == nil {
      continue
    }

    if !result {
      continue
    }

    fmt.Print("\t" + file.Name() + "\n")
  }
}

func SearchPerson(parameter []string) {

}

func ExportList(parameter []string) {

}

func main() {

  fmt.Print("\n### Go Anwesenheitsliste - Analyse Tool - begin ###\n\n")

  runCondition := true
  for runCondition {

    cmd, parameter := ReadStdIn()

    switch cmd {

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
      runCondition = false
    }

    fmt.Print("\n")
  }

  fmt.Print("### Go Anwesenheitsliste - Analyse Tool - end ###\n\n")
}
