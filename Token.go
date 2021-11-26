package main

import (
  "hash/fnv"
  "strconv"
  "time"
)

func creatAndRunTimer(interval int, locations []Location) {
  ticker := time.NewTicker(time.Duration(interval) * time.Second)
  quit := make(chan struct{})
  go func() {
    for {
      select {
      case <- ticker.C:
            runChangeTokenThread()

        case <- quit:
            ticker.Stop()
            return
        }
      }
  }()
}

func runChangeTokenThread() {
  locations, result := ReadLocationList()
  if !result {
    return
  }

  var newList []Location

  for _, location := range(locations) {
    var tokenStringToHash string
    timeStamp := time.Now()
    tokenStringToHash = location.Name
    tokenStringToHash = tokenStringToHash + string(rune(location.Id))
    tokenStringToHash = tokenStringToHash + timeStamp.String()

    hasher := fnv.New64a()
    hasher.Write([]byte(tokenStringToHash))

    hash := strconv.FormatUint(hasher.Sum64(), 10)

    newLocation := Location {
      Id: location.Id,
      Name: location.Name,
      AccessToken: location.AccessToken,
      CurrentToken: hash,
      OldToken: location.CurrentToken,
    }

    newList = append(newList, newLocation)
  }

  WriteLocationListToFile(newList)
}

// Mit dieser Funktion kann geprüft werden ob der Zeit Token in Kombination mit
// dem Zugang Token gültig ist
func isTokenValid(accessToken string, timeToken string) (bool){
  locations, result := ReadLocationList()
  if !result {
    return false
  }

  for _, location := range(locations) {
    if location.AccessToken == accessToken && (location.CurrentToken == timeToken || location.OldToken == timeToken) {
      return true
    }
  }

  return false
}
