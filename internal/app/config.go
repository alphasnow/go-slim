package app

import "path/filepath"

type Config struct {
	AppName     string `yaml:"app_name" env:"APP_NAME"`
	AppEnv      string `yaml:"app_env" env:"APP_ENV"`
	AppDebug    bool   `yaml:"app_debug" env:"APP_DEBUG"`
	AppUrl      string `yaml:"app_url" env:"APP_URL"`
	RootPath    string `yaml:"root_path" env:"ROOT_PATH"`
	LogsPath    string `yaml:"logs_path" env:"LOGS_PATH"`
	ConfigsPath string `yaml:"configs_path" env:"CONFIGS_PATH"`
	UploadsPath string `yaml:"uploads_path" env:"UPLOADS_PATH"`
	PrivatePath string `yaml:"private_path" env:"PRIVATE_PATH"`
	TmpPath     string `yaml:"tmp_path" env:"TMP_PATH"`
	CachePath   string `yaml:"cache_path" env:"CACHE_PATH"`
	BuildTag    string `yaml:"build_tag"`
	BuildDate   string `yaml:"build_date"`
}

func NewConfig() *Config {
	return &Config{
		AppName:  "APP",
		AppEnv:   Env,
		AppDebug: false,
		AppUrl:   "",

		BuildTag:  buildTag,
		BuildDate: buildDate,

		RootPath:    RootPath,
		LogsPath:    LogsPath,
		ConfigsPath: ConfigsPath,
		UploadsPath: UploadsPath,
		PrivatePath: PrivatePath,
		TmpPath:     TmpPath,
		CachePath:   CachePath,
	}
}

func (c *Config) JoinRoot(path string) string {
	return filepath.Join(c.RootPath, path)
}
func (c *Config) JoinUploads(path string) string {
	return filepath.Join(c.UploadsPath, path)
}
func (c *Config) JoinPrivate(path string) string {
	return filepath.Join(c.PrivatePath, path)
}
func (c *Config) IsDev() bool {
	return c.AppEnv == EnvDev
}
func (c *Config) IsProd() bool {
	return c.AppEnv == EnvProd
}
