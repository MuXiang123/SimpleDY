package initial

import (
	"SimpleDY/global"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//读取配置信息
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "viper readinconfig error: %v\n", err)
		os.Exit(1)

	}
	//反序列化到变量中
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Fprintf(os.Stderr, "viper unmarshal err: %v\n", err)
		os.Exit(1)
	}
}
