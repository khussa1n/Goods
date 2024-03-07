package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	project := router.Group("/projects")
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
