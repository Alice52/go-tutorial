package json

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Server  string `mapstructure:"server"`
	Port    int    `mapstructure:"port"`
	Enabled bool   `mapstructure:"enabled"`
}

func MapStructUsage() {
	configMap := map[string]interface{}{
		"server":  "example.com",
		"port":    8080,
		"enabled": true,
		"timeout": 30, // 该字段在结构体中没有对应的字段, 将被忽略
	}

	var conf Config
	err := mapstructure.Decode(configMap, &conf)
	if err != nil {
		fmt.Println("Error decoding config:", err)
		return
	}

	fmt.Println("Server:", conf.Server)
	fmt.Println("Port:", conf.Port)
	fmt.Println("Enabled:", conf.Enabled)
}
