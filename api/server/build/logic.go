package build

import (
	"github.com/google/wire"
	"go-slim/api/server/logic"
)

var LogicNewSet = wire.NewSet(
	wire.Struct(new(logic.UserClient), "*"),
)
