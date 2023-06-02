package xpaddle

type Config struct {
	Url string `yaml:"url" env:"PADDLE_URL"`
}

func NewConfig() *Config {
	return &Config{
		Url: "http://127.0.0.1:8866",
	}
}
