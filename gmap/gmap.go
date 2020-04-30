package gmap

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

//Location holds lat,long data
type Location struct {
	Name    string
	Country string
	Lat     float64
	Lng     float64
}

//Geocode returns location struct info for a given location ie. Chester, UK
func Geocode(address string, apiKey string) (loc Location) {
	var client *maps.Client
	var err error

	if apiKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(apiKey))
	}
	check(err)

	r := &maps.GeocodingRequest{
		Address: address,
	}

	resp, err := client.Geocode(context.Background(), r)
	check(err)
	return locationData(resp)

}

func locationData(resp []maps.GeocodingResult) (loc Location) {
	return Location{
		Lat:     resp[0].Geometry.Location.Lat,
		Lng:     resp[0].Geometry.Location.Lng,
		Name:    resp[0].AddressComponents[0].ShortName,
		Country: resp[0].AddressComponents[4].LongName,
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

}
