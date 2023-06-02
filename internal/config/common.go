package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func NewConfig(cfgPath string, envPath string) *Config {

	// new default
	cfg := DefaultConfig()

	// parse yaml
	parseYaml(cfg, cfgPath)

	// parse env
	parseEnv(cfg, envPath)

	log.Println("[APP_ENV]", cfg.App.AppEnv, cfg.App.AppDebug)
	log.Println("[APP_PATH]", cfg.App.RootPath, cfg.App.ConfigsPath, cfg.App.LogsPath, cfg.App.UploadsPath, cfg.App.PrivatePath, cfg.App.CachePath, cfg.App.TmpPath)
	log.Println("[APP_NAME]", cfg.App.AppName)

	return cfg
}

func parseYaml(cfg *Config, yamlDir string) {
	files, _ := ioutil.ReadDir(yamlDir)
	for _, file := range files {
		fName := file.Name()
		if file.IsDir() == true || filepath.Ext(fName) != ".yaml" {
			continue
		}
		fPath := filepath.Join(yamlDir, fName)
		// parse file
		cfgByte, err := ioutil.ReadFile(fPath)
		if err != nil {
			panic(err)
		}
		if err = yaml.Unmarshal(cfgByte, cfg); err != nil {
			panic(err)
		}
	}

}

func parseEnv(cfg *Config, envDir string) {
	// parse env
	envFile := filepath.Join(envDir, ".env")
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}
	if err = env.Parse(cfg); err != nil {
		panic(err)
	}
}
