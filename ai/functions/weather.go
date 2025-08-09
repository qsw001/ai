package functions

import(
	"fmt"
)

type WeatherParams struct {
	City string `json:"city"`
}

func GetWeather(params WeatherParams) (string, error) {
	//先模拟结果
	return fmt.Sprintf("%s今天天气晴，气温0度",params.City), nil
}