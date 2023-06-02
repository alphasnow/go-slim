package build

import (
	"github.com/google/wire"
	"go-slim/api/server/services/file_upload"
	"go-slim/api/server/services/sign_checker"
)

var ServicesNewSet = wire.NewSet(
	wire.Struct(new(sign_checker.SignJsonChecker), "*"),
	wire.Struct(new(sign_checker.SignUrlChecker), "*"),
	wire.Struct(new(sign_checker.SignChecker), "*"),
	wire.Struct(new(file_upload.FormFile), "*"),
	wire.Struct(new(file_upload.FormData), "*"),
)
