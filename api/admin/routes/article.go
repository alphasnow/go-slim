package routes

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/controllers"
)

func ArticleRouteRegister(
	r *gin.RouterGroup,
	ArticleContent *controllers.ArticleContentController,
	ArticleCategory *controllers.ArticleCategoryController,
) {
	{
		_r := r.Group("/article/content")
		_r.GET("/index", ArticleContent.Index)
		_r.GET("/create", ArticleContent.Create)
		_r.POST("/create", ArticleContent.Store)
		_r.GET("/show/:id", ArticleContent.Show)
		_r.GET("/edit/:id", ArticleContent.Edit)
		_r.POST("/edit/:id", ArticleContent.Update)
		_r.POST("/delete/:id", ArticleContent.Destroy)
		_r.GET("/form", ArticleContent.Form)
	}
	{
		_r := r.Group("/article/category")
		_r.GET("/index", ArticleCategory.Index)
		_r.GET("/create", ArticleCategory.Create)
		_r.POST("/create", ArticleCategory.Store)
		_r.GET("/show/:id", ArticleCategory.Show)
		_r.GET("/edit/:id", ArticleCategory.Edit)
		_r.POST("/edit/:id", ArticleCategory.Update)
		_r.POST("/delete/:id", ArticleCategory.Destroy)
		_r.GET("/form", ArticleCategory.Form)
	}
}
