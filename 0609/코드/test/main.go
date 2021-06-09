package main

import (
	"fmt"
	"src/geo"
)

func main() {
	location := geo.Landmark{}
	location.Name = "The GooglePlex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
