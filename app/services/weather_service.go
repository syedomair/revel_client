package services

import (
	"github.com/revel/revel"
)

type WeatherService struct {
	CommonService
}

func (c WeatherService) GetWeather(jsonStr []byte, session revel.Session) map[string]interface{} {
	return c.CallBackend("POST", "v1/weather", jsonStr, true, session)
}
