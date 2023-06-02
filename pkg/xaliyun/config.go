package xaliyun

type Config struct {
	AccessKeyId     string `yaml:"access_key_id" env:"ALIYUN_ACCESS_KEY_ID"`
	AccessKeySecret string `yaml:"access_key_secret" env:"ALIYUN_ACCESS_KEY_SECRET"`
	RegionId        string `yaml:"region_id" env:"ALIYUN_REGION_ID"`
}

func NewConfig() *Config {
	return &Config{
		RegionId: "cn-shanghai",
	}
}
