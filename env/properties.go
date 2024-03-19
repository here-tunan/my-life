package env

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Properties struct {
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mysql"`

	Redis struct {
		HOST     string `json:"host"`
		Port     string `json:"port"`
		Db       int    `json:"db"`
		Password string `json:"password"`
	} `json:"redis"`

	// 像这种 AaaBbb的 需要使用 yaml 标签对应配置文件中的内容才能被成功解析
	ElasticSearch struct {
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Fingerprint string `json:"fingerprint"`
	} `json:"elasticSearch" yaml:"elasticSearch"`
}

var Prop *Properties

// 初始化配置文件
func init() {
	readEncProperties()
}

func readEncProperties() {

	envFilePath := "./env/dev.yaml"

	// 根据环境变量解析配置文件
	appEnv := os.Getenv("GO_MY_LIFE_ENV")
	if appEnv == "dev" || appEnv == "" {
		fmt.Println("Start go-my-life app in development environment!")
		envFilePath = "./env/dev.yaml"
	} else if appEnv == "prod" {
		fmt.Println("Start go-my-life app in production environment!")
		envFilePath = "./env/prod.yaml"
	} else {
		fmt.Println("Start go-my-life app in unknown environment! Maybe cause errors!")
	}

	// 读取YAML文件内容
	envFile, err := os.ReadFile(envFilePath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	prop := &Properties{}

	// 解析YAML文件
	err = yaml.Unmarshal(envFile, prop)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
	// 打印解析后的数据
	// fmt.Printf("%+v\n", prop)

	Prop = prop
}
