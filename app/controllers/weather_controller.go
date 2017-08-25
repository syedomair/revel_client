package controllers

import (
	"github.com/revel/revel"
	"github.com/syedomair/revel_client/app/services"
)

type WeatherController struct {
	*revel.Controller
}

func (c WeatherController) GetWeather() revel.Result {

	return c.Render()
}

func (c WeatherController) WeatherData(address string) revel.Result {
	strJson := "{\"address\":\"" + address + "\" }"

	c.Session["server"] = "weather"
	resultMap := services.WeatherService{}.GetWeather([]byte(strJson), c.Session)
	response := "error"
	if resultMap["result"] == "success" {
		resultMap1 := make(map[string]interface{})
		resultMap1 = resultMap["data"].(map[string]interface{})

		address := resultMap1["address"]
		temperature := resultMap1["temperature"]
		response = "{\"address\":\"" + address.(string) + "\", \"temperature\":\"" + temperature.(string) + "\"}"
	}
	return c.RenderJSON(response)
}
