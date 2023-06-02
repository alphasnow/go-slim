package xcrypt

type Config struct {
	Secret string `yaml:"secret"`
	Iv     string `yaml:"iv"`
}

func NewConfig() *Config {
	return &Config{
		Secret: "#Secret##Secret#",
		Iv:     "#Iv##Iv##Iv##Iv#",
	}
}

func (c *Config) GetSecret() []byte {
	return []byte(c.Secret)
}
func (c *Config) GetIv() []byte {
	return []byte(c.Iv)
}
