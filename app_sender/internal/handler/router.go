package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/khussa1n/Goods/app_sender/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	project := router.Group("/project")
	project.POST("/create", h.createProject)
	project.PATCH("/update", h.updateProjectByID)
	project.DELETE("/remove", h.deleteProjectByID)

	projects := router.Group("/projects")
	projects.GET("/list", h.getAllProjects)

	good := router.Group("/good")
	good.POST("/create", h.createGood)
	good.PATCH("/update", h.updateGoodByID)
	good.DELETE("/remove", h.deleteGoodByID)
	good.PATCH("/reprioritiize", h.reprioritiize)

	goods := router.Group("/goods")
	goods.GET("/list", h.getAllGoods)

	return router
}
