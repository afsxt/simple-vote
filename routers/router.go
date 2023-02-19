package routers

import (
	"github.com/afsxt/simple-vote/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	adminV1 "github.com/afsxt/simple-vote/routers/admin/v1"
	apiV1 "github.com/afsxt/simple-vote/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// todo: 暂没考虑用户这一套逻辑，不进行用户认证
	adminGroupV1 := r.Group("/admin/v1")
	// adminGroupV1.Use(jwt.JWT())
	{
		adminGroupV1.POST("/vote/theme", adminV1.AddTheme)
		adminGroupV1.PUT("/vote/theme/:id", adminV1.UpdateTheme)
		adminGroupV1.POST("/vote/candidates", adminV1.AddCandidates)
		adminGroupV1.POST("/vote/theme/:id/state", adminV1.ChangeThemeState)
		adminGroupV1.GET("/vote/theme/:themeID", adminV1.GetThemeResult)
		adminGroupV1.GET("/vote/theme/:themeID/candidate/:candidateID/users", adminV1.GetCandidateUsers)
	}

	apiGroupv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	{
		apiGroupv1.POST("/vote/verify", apiV1.VerifyUser)
		apiGroupv1.POST("/vote", apiV1.Vote)
		apiGroupv1.GET("/vote/theme/:themeID", apiV1.GetVoteDetails)
	}

	return r
}
