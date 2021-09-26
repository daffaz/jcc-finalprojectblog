package Routes

import (
	"jccblog/Controllers"
	"jccblog/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// docs "jccblog/docs"

// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"

var adminList = map[string]string{}
var writerList = map[string]string{}

func SetupRouter() *gin.Engine {
	Utils.GetAdminAccount()
	Utils.GetWriterAccount()
	r := gin.Default()

	for _, list := range Utils.Admin {
		adminList[list.Username] = list.Password
	}

	for _, list := range Utils.Writer {
		writerList[list.Username] = list.Password
	}

	if len(adminList) == 0 {
		adminList["admin"] = "admin"
	}
	if len(writerList) == 0 {
		writerList["dummy"] = "dummy"
	}

	// Auth for admin, case like fetch all user data, delete user data if needed and etc
	adminApiGroup := r.Group("/api", gin.BasicAuth(adminList))

	// Auth for casual user, where the role is writer
	apiGroup := r.Group("/api", gin.BasicAuth(writerList))

	{
		adminApiGroup.GET("users", Controllers.GetAllUsers)
		adminApiGroup.GET("users/:id", Controllers.GetUserById)
		adminApiGroup.GET("bantuan", Controllers.GetAllBantuan)
		adminApiGroup.POST("categories", Controllers.CreateNewCategory)
		adminApiGroup.DELETE("users", Controllers.DeleteUser)
	}
	{
		/*
		 * ----- User
		**/
		apiGroup.PUT("users", Controllers.UpdateUser)
		/*
		 * ----- UserMeta
		**/
		apiGroup.GET("usermeta/:id", Controllers.GetUserMeta)
		apiGroup.POST("usermeta/", Controllers.CreateNewUserMeta)
		/*
		 * ----- Blogs
		**/
		apiGroup.GET("blogs", Controllers.GetAllBlogs)
		apiGroup.GET("blogs/:id", Controllers.GetBlogDataById)
		apiGroup.GET("blogs/title/:slug", Controllers.GetBlogDataBySlug)
		apiGroup.PUT("blogs/:id", Controllers.UpdateBlog)
		apiGroup.POST("blogs", Controllers.CreateNewBlog)
		apiGroup.DELETE("blogs/:id", Controllers.DeleteBlog)
		/*
		 * ----- Categories
		**/
		apiGroup.GET("categories", Controllers.GetAllCategories)
		/*
		 * ----- Comments
		**/
		apiGroup.GET("comment/:id", Controllers.GetCommentByBlogID)
		apiGroup.POST("comment", Controllers.CreateNewComment)
		apiGroup.DELETE("comment", Controllers.DeleteComment)
		/*
		 * ----- Bantuan
		**/
		apiGroup.POST("bantuan", Controllers.CreateNewbantuan)
		/*
		 * ----- Favorit
		**/
		apiGroup.GET("favorit/user_id/:id", Controllers.GetAllFavoritByUID)
		apiGroup.POST("favorit", Controllers.CreateNewFavorit)
	}
	// This is for the new user, so they can register
	r.POST("api/register", Controllers.CreateNewUser)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello you found the root of the this API URL, go to http://localhost:8080/docs/index.html for detailed documentation what you can do")
	})

	// docs.SwaggerInfo.BasePath = "/"
	// r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
