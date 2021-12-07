package main

import (
  "hash/fnv"
  "strconv"
  "time"
)

// This function start a ticker thread, which change the time token in a definite interval
func creatAndRunTimer(interval int, locations []Location) {
  // Create ticker
  ticker := time.NewTicker(time.Duration(interval) * time.Second)

  // Start ticker
  quit := make(chan struct{})
  go func() {
    for {
      select {
      case <- ticker.C:
            // Run function to change the time token
            runChangeTokenThread()

        case <- quit:
            ticker.Stop()
            return
        }
      }
  }()
}

// This function change the time token of each location
func runChangeTokenThread() {
  // Get the locations
  locations, result := ReadLocationList()
  if !result {
    return
  }

  var newList []Location

  // Iterate throw all locations
  for _, location := range(locations) {
    // Create new token
    var tokenStringToHash string
    timeStamp := time.Now()
    tokenStringToHash = location.Name
    tokenStringToHash = tokenStringToHash + string(rune(location.Id))
    tokenStringToHash = tokenStringToHash + timeStamp.String()

    hasher := fnv.New64a()
    hasher.Write([]byte(tokenStringToHash))

    hash := strconv.FormatUint(hasher.Sum64(), 10)

    // Set new token into location
    newLocation := Location {
      Id: location.Id,
      Name: location.Name,
      AccessToken: location.AccessToken,
      CurrentToken: hash,
      OldToken: location.CurrentToken,
    }

    // Add modified location into list
    newList = append(newList, newLocation)
  }

  // Write locations with new tokens into XML file
  WriteLocationListToFile(newList)
}

// This function validate with time and access token.
func isTokenValid(accessToken string, timeToken string) (bool){
  // Reading configuration file
  locations, result := ReadLocationList()
  if !result {
    return false
  }

  // Iterating throw all locations
  for _, location := range(locations) {
    // Checking time and access token
    if location.AccessToken == accessToken && (location.CurrentToken == timeToken || location.OldToken == timeToken) {
      return true
    }
  }

  return false
}
