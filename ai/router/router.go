package router

import (
	"ai/handler"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/chat", handler.ChatHandler)
}