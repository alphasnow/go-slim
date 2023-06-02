package xmysql

import (
	"database/sql"
	"fmt"
	"go-slim/pkg/xlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

func NewMysql(cfg *Config, logManager *xlog.Manager) *gorm.DB {
	if err := pingDatabase(cfg.Write); err != nil {
		panic(err)
	}

	if cfg.Base.EnableCreateDatabase {
		if err := createDatabase(cfg.Write); err != nil {
			panic(err)
		}
	}

	// https://gorm.io/docs/logger.html
	zLog := logManager.Logger(xlog.SQL)
	xWriter := &Writer{Zap: zLog}
	log := logger.New(xWriter, logger.Config{
		SlowThreshold:             cfg.Logger.GetSlowThreshold(), // Slow SQL threshold
		LogLevel:                  cfg.Logger.GetLogLevel(),      // Log level
		IgnoreRecordNotFoundError: true,                          // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,                         // Disable color
	})

	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dialector := mysql.New(mysql.Config{
		DSN: cfg.Write.DNS(),
	})
	gcfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Base.TablePrefix,
			SingularTable: cfg.Base.SingularTable,
		},
		// https://gorm.io/zh_CN/docs/performance.html
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 log,
	}
	db, err := gorm.Open(dialector, gcfg)
	if err != nil {
		panic(err)
	}

	if cfg.Base.Debug {
		db = db.Debug()
	}

	if len(cfg.Reads) > 0 {
		var replicas []gorm.Dialector
		for _, read := range cfg.Reads {
			replicas = append(replicas, mysql.Open(read.DNS()))
		}
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(cfg.Base.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Base.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Base.MaxLifetime) * time.Second)

	return db
}

func pingDatabase(c DSNConfig) error {
	db, err := sql.Open("mysql", c.DB())
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	return err
}

func createDatabase(cfg DSNConfig) error {
	db, err := sql.Open("mysql", cfg.DB())
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.Database)
	_, err = db.Exec(query)
	return err
}
