package xlimiter

type Config struct {
	DefaultRate   string `yaml:"defaultRate"`
	LimiterPrefix string `yaml:"limiterPrefix"`
	RedisPrefix   string `yaml:"redisPrefix"`
}

func (c *Config) JoinPrefix(prefix string) string {
	if c.RedisPrefix == "" {
		return c.LimiterPrefix + ":" + prefix
	}
	return c.RedisPrefix + ":" + c.LimiterPrefix + ":" + prefix
}

func NewConfig() *Config {
	return &Config{
		DefaultRate:   "1-S",
		LimiterPrefix: "x-ratelimit",
		RedisPrefix:   "",
	}
}
