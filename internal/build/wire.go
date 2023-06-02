//go:build wireinject
// +build wireinject

package build

import (
	"github.com/google/wire"
	"go-slim/internal/config"
)

func BuildApp(cfg *config.Config) (*App, error) {

	wire.Build(
		wire.Struct(new(App), "*"),
		InternalNewSet,
		ModulesNewSet,
		PkgNewSet,
	)
	return &App{}, nil
}

// BuildCli
func BuildCli(cfg *config.Config) *Cli {

	wire.Build(
		wire.Struct(new(Cli), "*"),
		InternalNewSet,
		PkgNewSet,
	)
	return &Cli{}
}
