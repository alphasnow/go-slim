package app

import (
	"flag"
	"log"
	"path/filepath"
)

const EnvDev = "dev"
const EnvProd = "prod"
const EnvTest = "test"

var (
	Env   string
	Debug bool

	RootPath    string
	LogsPath    string
	ConfigsPath string
	UploadsPath string
	PrivatePath string
	CachePath   string
	TmpPath     string
)

func init() {
	flag.StringVar(&Env, "env", EnvDev, "set app env")
	flag.BoolVar(&Debug, "debug", false, "set app debug")

	flag.StringVar(&RootPath, "root_path", "", "set root path")
	flag.StringVar(&LogsPath, "logs_path", "", "set logs path")
	flag.StringVar(&ConfigsPath, "configs_path", "", "set configs path")
	flag.StringVar(&UploadsPath, "uploads_path", "", "set upload path")
	flag.StringVar(&PrivatePath, "private_path", "", "set private file path")
	flag.StringVar(&CachePath, "cache_path", "", "set file cache path")
	flag.StringVar(&TmpPath, "tmp_path", "", "set file tmp path")
	flag.Parse()

	log.Println("[FLAG_ENV]", Env, Debug)
	log.Println("[FLAG_PATH]", RootPath, ConfigsPath, LogsPath, UploadsPath, PrivatePath, CachePath, TmpPath)

	autoSetPath()
}

func autoSetPath() {

	if RootPath == "" {
		if Env == EnvProd {
			// release
			RootPath = ExecRootPath()
		} else {
			// cmd/server/main.go
			RootPath = CallRootPath()
		}
	}

	if ConfigsPath == "" {
		ConfigsPath = filepath.Join(RootPath, "configs")
	}

	if LogsPath == "" {
		LogsPath = filepath.Join(RootPath, "storage/logs")
	}
	if UploadsPath == "" {
		UploadsPath = filepath.Join(RootPath, "storage/uploads")
	}
	if PrivatePath == "" {
		PrivatePath = filepath.Join(RootPath, "storage/private")
	}
	if CachePath == "" {
		CachePath = filepath.Join(RootPath, "storage/cache")
	}
	if TmpPath == "" {
		TmpPath = filepath.Join(RootPath, "storage/tmp")
	}

	log.Println("[INIT_PATH]", RootPath, ConfigsPath, LogsPath, UploadsPath, PrivatePath, CachePath, TmpPath)
}
