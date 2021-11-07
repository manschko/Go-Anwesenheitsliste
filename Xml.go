package main

type Locations struct {
	Location []Location `xml:"location"`
}

type Location struct {
	Name string `xml:"name"`
}
