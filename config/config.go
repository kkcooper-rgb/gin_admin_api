package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// config 总配置文件
type config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
}

// 系统配置
type system struct {
	Host string `yaml:"host"`
	Post int    `ymal:"post"`
	Env  string `ymal:"env"`
}

type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

var Config *config

func init() {
	file, err := os.Open("./config.yaml")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	yaml.Unmarshal([]byte(content), &Config)
}
