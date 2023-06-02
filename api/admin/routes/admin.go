package routes

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/controllers"
)

func AdminRouteRegister(
	r *gin.RouterGroup,
	AdminUser *controllers.AdminUserController,
	AdminDepartment *controllers.AdminDepartmentController,
) {
	{
		_r := r.Group("/admin/user")
		_r.GET("/index", AdminUser.Index)
		_r.GET("/create", AdminUser.Create)
		_r.POST("/create", AdminUser.Store)
		_r.GET("/show/:id", AdminUser.Show)
		_r.GET("/edit/:id", AdminUser.Edit)
		_r.POST("/edit/:id", AdminUser.Update)
		_r.POST("/delete/:id", AdminUser.Destroy)
		_r.GET("/form", AdminUser.Form)
	}
	{
		_r := r.Group("/admin/department")
		_r.GET("/index", AdminDepartment.Index)
		_r.GET("/create", AdminDepartment.Create)
		_r.POST("/create", AdminDepartment.Store)
		_r.GET("/show/:id", AdminDepartment.Show)
		_r.GET("/edit/:id", AdminDepartment.Edit)
		_r.POST("/edit/:id", AdminDepartment.Update)
		_r.POST("/delete/:id", AdminDepartment.Destroy)
		_r.GET("/form", AdminDepartment.Form)
	}
}
