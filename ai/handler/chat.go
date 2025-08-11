package handler

import (
	"ai/db"
	"ai/functions"
	"ai/model"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cloudwego/eino/schema"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "无效请求", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	chatModel, err := model.NewArkModel(ctx)
	if err != nil {
		http.Error(w, "模型初始化失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	history, err := db.GetChatHistory()
	if err != nil {
		http.Error(w, "加载历史失败", http.StatusInternalServerError)
		return
	}

	messages := []*schema.Message{
		schema.SystemMessage("你是可爱的女高中生"),
	}

	for _, msg := range history {
		messages = append(messages, schema.UserMessage(msg))
	}

	userInput := req.Message
	messages = append(messages, schema.UserMessage(userInput))
	db.SaveMessage(userInput)

	// 简单检测是否需要调用函数
	if strings.Contains(userInput, "天气") {
    city := extractCityFromText(userInput)

    weatherResult, err := functions.GetWeather(functions.WeatherParams{City: city})
    if err != nil {
        weatherResult = "获取天气时出错了"
    }

    messages = append(messages, schema.UserMessage(userInput))
    messages = append(messages, schema.AssistantMessage(weatherResult, nil))

    _ = db.SaveMessage(userInput)
    _ = db.SaveMessage(weatherResult)

    // return weatherResult // 根据你的需求决定是否返回
}

	stream, err := chatModel.Stream(ctx, messages)
	if err != nil {
		http.Error(w, "模型调用失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stream.Close()

	var reply strings.Builder
	for {
		chunk, err := stream.Recv()
		if err != nil {
			break
		}
		reply.WriteString(chunk.Content)
	}
	answer := reply.String()
	db.SaveMessage(answer)

	resp := ChatResponse{Reply: answer}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// 简单提取城市的函数示例
func extractCityFromText(text string) string {
	// 假设用户说“请告诉我上海的天气”
	if strings.Contains(text, "上海") {
		return "上海"
	}
	if strings.Contains(text, "北京") {
		return "北京"
	}
	// 其他城市可加
	return ""
}