package main

import (
	"go-my-life/api"
	_ "go-my-life/internal/infrastructure"
	"log"
)

func main() {
	log.Println("Go my life is starting!")
	// 启用web服务
	api.Start()
}
