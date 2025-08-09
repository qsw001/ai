// package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/cloudwego/eino-ext/components/model/ark"
// 	"github.com/cloudwego/eino/schema"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// 1. 加载环境变量，包含 ARK_API_KEY 和 MODEL
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		panic("加载 .env 文件失败: " + err.Error())
// 	}

// 	ctx := context.Background()

// 	// 2. 创建豆包聊天模型实例，传入密钥和模型名
// 	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
// 		APIKey: os.Getenv("ARK_API_KEY"),
// 		Model:  os.Getenv("MODEL"), 
// 	})
// 	if err != nil {
// 		panic("创建模型失败: " + err.Error())
// 	}

// 	// 3. 创建对话历史消息数组，先加一个系统角色设定
// 	messages := []*schema.Message{
// 		schema.SystemMessage("你是一个AI助手"),
// 	}

// 	// 4. 读取用户输入，循环对话
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 			fmt.Print("你：")
// 			userInput, _ := reader.ReadString('\n')
// 			userInput = strings.TrimSpace(userInput)
// 			if userInput == "退出" || userInput == "exit" {
// 				fmt.Println("对话结束，再见！")
// 				break
// 			}

// 		// 5. 用户输入追加到对话历史
// 		messages = append(messages, schema.UserMessage(userInput))

// 				// 6. 使用流式调用模型生成回答
// 		stream, err := model.Stream(ctx, messages)
// 		if err != nil {
// 			fmt.Println("模型调用失败:", err)
// 			continue
// 			}
// 		defer stream.Close()

// 		var contentBuilder strings.Builder

// 		for {
// 			chunk, err := stream.Recv()
// 			if err != nil {
// 				// 流结束或出错，跳出循环
// 				break
// 			}
// 			fmt.Print(chunk.Content)
// 			contentBuilder.WriteString(chunk.Content)
// 		}
// 		fmt.Println()

// 		// 7. 模型回答追加到历史
// 		messages = append(messages, schema.AssistantMessage(contentBuilder.String(), nil))
// 	}
// }


//向量化过程
// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/cloudwego/eino-ext/components/embedding/ark"
// 	_"github.com/cloudwego/eino/components/embedding"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load(".env")
// 	if err != nil{
// 		panic(err)
// 	}
// 	ctx := context.Background()

// 	embedder,err := ark.NewEmbedder(ctx,&ark.EmbeddingConfig{
// 		APIKey:	os.Getenv("ARK_API_KEY"),
// 		Model: os.Getenv("EMBEDDER"),
// 	})

// 	if err!=nil{
// 		panic(err)
// 	}
	
// 	input:=[]string{
// 		"你好,泥豪",
// 	}

// 	embeddings, err := embedder.EmbedStrings(ctx,input)
// 	if err!=nil{
// 		panic(err)
// 	}

// 	fmt.Println(embeddings)
// }

package main

import (
	"ai/config"
	"ai/db"
	"ai/router"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	db.InitMySQL()

	router.RegisterRoutes()

	log.Println(" 服务启动于 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}