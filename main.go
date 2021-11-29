package main

func main() {
	ConfigFlag()

	locations, result := ReadLocationList()
	if !result {
		return
	}

	creatAndRunTimer(flags.TokenValidity, locations)

	ConfigWebServer()
	wg.Wait()
}
