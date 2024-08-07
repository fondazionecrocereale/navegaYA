package router

import (
	"fmt"

	"github.com/fondazionecrocereale/navegayaapi/controller"

	"net/http"

	"github.com/fondazionecrocereale/navegayaapi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 	Tag Service API
func NewRouter(tagsController *controller.TagsController, clipsController *controller.ClipsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	fmt.Print(docs.SwaggerInfo.BasePath)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	clipsRouter := baseRouter.Group("/clips")
	clipsRouter.GET("", clipsController.FindAll)
	clipsRouter.GET("/:clipId", clipsController.FindById)
	clipsRouter.POST("", clipsController.Create)
	clipsRouter.PATCH("/:clipId", clipsController.Update)
	clipsRouter.DELETE("/:clipId", clipsController.Delete)

	return router
}
