package getWkuTemp

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
);

type WeatherData struct {
	ZBarometer              string `json:"ZBarometer"`
	BarometerTendency       string `json:"BarometerTendency"`
	Thermometer             string `json:"Thermometer"`
	Dewpoint                string `json:"Dewpoint"`
	HeatIndex               string `json:"HeatIndex"`
	WetBulbGlobeTemperature string `json:"WetBulbGlobeTemperature"`
	WindChill               string `json:"WindChill"`
	Anemometer              string `json:"Anemometer"`
	One0MinuteWindGust      string `json:"10MinuteWindGust"`
	Hygrometer              string `json:"Hygrometer"`
	WindVane                string `json:"WindVane"`
	SolarRadiationSensor    string `json:"SolarRadiationSensor"`
	UVRadiationSensor       string `json:"UVRadiationSensor"`
	RainRate                string `json:"RainRate"`
	RainGauge               string `json:"RainGauge"`
	string                  `json:""`
	CondensationDew         string `json:"CondensationDew"`
	ObTime                  string `json:"obTime"`
}

func GetTemp () string {

	// Make the request
	res, err := http.Get("http://wkuweather.com/data/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Close the response body
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Read the response body
	doc, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode Json into an interface
	var weatherData WeatherData
	b := []byte(string(doc))
	err = json.Unmarshal(b, &weatherData)
	if err != nil {
		log.Fatal(err)
	}

	return weatherData.WindChill
}
