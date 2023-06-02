package xsession

type Config struct {
	Name   string       `yaml:"name" env:"SESSION_NAME"`
	Store  string       `yaml:"store" env:"SESSION_STORE"`
	Cookie ConfigCookie `yaml:"cookie"`
	//Redis  ConfigRedis  `yaml:"redis"`
}
type ConfigCookie struct {
	Secret string `yaml:"secret" env:"SESSION_SECRET"`
}

//type ConfigRedis struct {
//	Host string `yaml:"host" env:"REDIS_HOST"`
//	Port int    `yaml:"port" env:"REDIS_PORT"`
//	Pass string `yaml:"pass" env:"REDIS_PASS"`
//	DB   int    `yaml:"db" env:"REDIS_DB"`
//}

func NewConfig() *Config {
	return &Config{
		Name:  "_session",
		Store: "cookie",
		Cookie: ConfigCookie{
			Secret: "#Secret#",
		},
		//Redis: ConfigRedis{
		//	Host: "127.0.0.1",
		//	Port: 6379,
		//	Pass: "",
		//	DB:   0,
		//},
	}
}
