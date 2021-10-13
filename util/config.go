package util

import (
	"github.com/spf13/viper"
	"time"
)

// Config store all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variable.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // 设置配置文件路径
	viper.SetConfigName("app") // 设置配置文件名称  `不包含拓展名`
	viper.SetConfigType("env") // 设置配置文件类型，可以是json、xml

	viper.AutomaticEnv() // 读取环境变量的值，并且用config中的值覆盖
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
