package xmysql

import (
	"go-slim/internal/app"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// NewGenerator
// https://github.com/go-gorm/gen
func NewGenerator(cfg *Config, global *app.Config) *gen.Generator {

	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath:      global.JoinRoot("internal/dal/query"),
		ModelPkgPath: global.JoinRoot("internal/dal/model"),
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// if you want generate field with unsigned integer type, set FieldSignable true
		FieldSignable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		//if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	db, _ := gorm.Open(mysql.Open(cfg.Write.DNS()))
	g.UseDB(db)
	return g
}
