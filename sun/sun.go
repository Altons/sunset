package sun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Time holds json respond
type Time struct {
	Sunrise                   string `json:"sunrise"`
	Sunset                    string `json:"sunset"`
	SolarNoon                 string `json:"solar_noon"`
	Daylength                 string `json:"day_length"`
	CivilTwilightBegin        string `json:"civil_twilight_begin"`
	CivilTwilightEnd          string `json:"civil_twilight_end"`
	NauticalTwilightBegin     string `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       string `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
}

type result map[string]interface{}

//GetSunData retrieve json respond from api
func GetSunData(lat float64, lng float64) (t Time) {
	url := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%g&lng=%g", lat, lng)
	req, err := http.Get(url)
	check(err)
	bodybytes, err := ioutil.ReadAll(req.Body)
	check(err)

	var res result
	json.Unmarshal(bodybytes, &res)

	itemList := res["results"].(map[string]interface{})
	var ret Time
	// convert map to json
	jsonString, _ := json.Marshal(itemList)
	//fmt.Println(string(jsonString))

	// convert json to struct
	json.Unmarshal(jsonString, &ret)
	//fmt.Println(ret)

	return ret
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
