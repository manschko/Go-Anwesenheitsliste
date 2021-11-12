package main

import (
	"fmt"
)

func main() {
	ConfigFlag()

	locations, result := ReadLocationList()
	if !result {
		return
	}

	for _, location := range locations {
		fmt.Printf("Id: %4v Name: %v\n", location.Id, location.Name)
	}

	creatAndRunTimer(flags.TokenValidity, locations)

	ConfigWebServer()
	wg.Wait()
}
