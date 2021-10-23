package main

func main() {
	ConfigFlag()
	ConfigWebServer()

	placeList, result := ReadLocationList()
	if !result {
		return
	}

	for _, location := range(placeList) {
		
	}
}
