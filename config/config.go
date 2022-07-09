package config

//服务器全局配置
type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
}

//Mysql数据库配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

// LogConfig 日志实体类
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}
