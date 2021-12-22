package main

/*
Matrikelnummern:
3186523
9008480
6196929
*/
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
