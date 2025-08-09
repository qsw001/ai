package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("加载环境变量失败：%v",err)
	}
}

func Getenv(key string) string{
	return os.Getenv(key)
}