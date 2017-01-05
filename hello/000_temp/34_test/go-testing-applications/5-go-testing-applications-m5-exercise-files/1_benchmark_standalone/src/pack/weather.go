package pack

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

func PrintWeather(zipCode int) {
	resp, err := http.Get("http://wsf.cdyne.com//WeatherWS/" +
		"Weather.asmx/GetCityWeatherByZIP?ZIP=" + strconv.Itoa(zipCode))

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)

	var weather Weather

	err = xml.Unmarshal(data, &weather)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(weather)
}

type Weather struct {
	State            string
	City             string
	Temperature      float32
	RelativeHumidity float32
	Wind             string
	Pressure         string
}
