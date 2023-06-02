package routes

import (
	"github.com/google/wire"
	"go-slim/api/admin/controllers"
)

type Controllers struct {
	Auth            *controllers.AuthController
	AdminUser       *controllers.AdminUserController
	AdminDepartment *controllers.AdminDepartmentController
	ArticleContent  *controllers.ArticleContentController
	ArticleCategory *controllers.ArticleCategoryController
	Server          *controllers.ServerController
}

var ControllersNewSet = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
	wire.Struct(new(controllers.AuthController), "*"),
	wire.Struct(new(controllers.AdminUserController), "*"),
	wire.Struct(new(controllers.AdminDepartmentController), "*"),
	wire.Struct(new(controllers.ArticleContentController), "*"),
	wire.Struct(new(controllers.ArticleCategoryController), "*"),
	wire.Struct(new(controllers.ServerController), "*"),
)
