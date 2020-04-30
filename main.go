package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/altons/sunset/gmap"
	"github.com/altons/sunset/sun"
)

type info struct {
	location gmap.Location
	times    sun.Time
}

func main() {
	apiKey := os.Getenv("GMAP_API_KEY") //Google Map API key
	address := flag.String("address", "Chester, UK", "address you want to check sunset and sunrise times")
	flag.Parse()

	loc := gmap.Geocode(*address, apiKey)
	times := sun.GetSunData(loc.Lat, loc.Lng)
	var data info
	data.location = loc
	data.times = times
	fmt.Printf("Sunrise and Sunset information for %s, %s\n", data.location.Name, data.location.Country)
	fmt.Println("Sunrise:", data.times.Sunrise)
	fmt.Println("Sunset:", data.times.Sunset)
	fmt.Println("Day Length:", data.times.Daylength)

}
