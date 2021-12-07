package main

func main() {
	ConfigFlag()

	locations, result := ReadLocationList()
	if !result {
		return
	}

	CreateAndRunTimer(flags.TokenValidity, locations)

	ConfigWebServer()
	wg.Wait()
}
