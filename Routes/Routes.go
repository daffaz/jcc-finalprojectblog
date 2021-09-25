package Routes

import (
	"jccblog/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api", gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
	}))
	{
		apiGroup.GET("blogs", Controllers.GetAllBlogs)
		apiGroup.GET("blogs/:id", Controllers.GetBlogDataById)
		apiGroup.GET("blogs/title/:slug", Controllers.GetBlogDataBySlug)
		/*
		 * ----- Categories
		**/
		apiGroup.GET("categories", Controllers.GetAllCategories)
		apiGroup.POST("categories", Controllers.CreateNewCategory)
		/*
		 * ----- User
		**/
		apiGroup.GET("users", Controllers.GetAllUsers)
	}
	r.POST("api/register", Controllers.CreateNewUser)

	return r
}
