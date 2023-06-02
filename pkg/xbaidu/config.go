package xbaidu

type Config struct {
	ApiKey    string `yaml:"api_key" env:"BAIDU_API_KEY"`
	SecretKey string `yaml:"secret_key" env:"BAIDU_SECRET_KEY"`
	AppID     int    `yaml:"app_id" env:"BAIDU_APP_ID"`
}

func NewConfig() *Config {
	return &Config{}
}
