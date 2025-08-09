package model

import (
	"context"
	"encoding/json"
	"os"

	"ai/functions" // 你需要创建这个目录
	"github.com/cloudwego/eino-ext/components/model/ark"
	_"github.com/cloudwego/eino/schema"
)

func NewArkModel(ctx context.Context) (*ark.ChatModel, error) {
	return ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
}
// 执行函数调用
func HandleFunctionCall(name string, args json.RawMessage) (string, error) {
	switch name {
	case "getWeather":
		var p functions.WeatherParams
		if err := json.Unmarshal(args, &p); err != nil {
			return "", err
		}
		return functions.GetWeather(p)
	default:
		return "未知函数", nil
	}
}