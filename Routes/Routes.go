package Routes

import (
	"jccblog/Controllers"
	"jccblog/Utils"
	"log"

	"github.com/gin-gonic/gin"
)

var adminList = map[string]string{}
var writerList = map[string]string{}

func SetupRouter() *gin.Engine {
	Utils.GetAdminAccount()
	r := gin.Default()

	for _, list := range Utils.Admin {
		adminList[list.Username] = list.Password
	}

	for _, list := range Utils.Writer {
		writerList[list.Username] = list.Password
	}

	log.Println(adminList)
	// Auth for admin, case like fetch all user data, delete user data if needed and etc
	adminApiGroup := r.Group("/api", gin.BasicAuth(adminList))

	// Auth for casual user, where the role is writer
	apiGroup := r.Group("/api", gin.BasicAuth(writerList))

	{
		adminApiGroup.GET("users", Controllers.GetAllUsers)
		adminApiGroup.DELETE("users", Controllers.GetAllUsers)
	}
	{
		/*
		 * ----- Blogs
		**/
		apiGroup.GET("blogs", Controllers.GetAllBlogs)
		apiGroup.POST("blogs", Controllers.CreateNewBlog)
		apiGroup.GET("blogs/:id", Controllers.GetBlogDataById)
		apiGroup.GET("blogs/title/:slug", Controllers.GetBlogDataBySlug)
		/*
		 * ----- Categories
		**/
		apiGroup.GET("categories", Controllers.GetAllCategories)
		apiGroup.POST("categories", Controllers.CreateNewCategory)
	}
	r.POST("api/register", Controllers.CreateNewUser)

	return r
}
