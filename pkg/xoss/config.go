package xoss

type Config struct {
	AccessKeyID     string `yaml:"access_key_id" env:"OSS_ACCESS_KEY_ID"`
	AccessKeySecret string `yaml:"access_key_secret" env:"OSS_ACCESS_KEY_SECRET"`
	Endpoint        string `yaml:"endpoint" env:"OSS_ENDPOINT"`
	Bucket          string `yaml:"bucket" env:"OSS_BUCKET"`
}

func NewConfig() *Config {
	return &Config{Endpoint: "oss-cn-shanghai.aliyuncs.com"}
}
