package rotation

import "path/filepath"

type Config struct {
	Filename   string `json:"filename" yaml:"filename" env:"LOG_FILENAME"`
	Filepath   string `json:"filepath" yaml:"filepath" env:"LOG_FILEPATH"`
	MaxSize    int    `json:"maxsize" yaml:"maxsize" env:"LOG_MAX_SIZE"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool   `json:"localtime" yaml:"localtime"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

func (c *Config) File() string {
	return filepath.Join(c.Filepath, c.Filename)
}
