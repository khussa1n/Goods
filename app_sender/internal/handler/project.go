package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/Goods/app_sender/internal/custom_error"
	"github.com/khussa1n/Goods/app_sender/internal/entity"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createProject(ctx *gin.Context) {
	var req entity.Projects
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 1, Message: custom_error.ErrInvalidInputBody.Error()})
		return
	}

	project, err := h.srvs.CreateProject(ctx, &req)
	if err != nil {
		log.Printf("can not create project: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 2, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, project)
}

func (h *Handler) getAllProjects(ctx *gin.Context) {
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.Error{Code: 1, Message: "invalid id param"})
		return
	}

	projectsList, err := h.srvs.GetAllProjects(ctx, limit, offset)
	if err != nil {
		log.Printf("can not get all projects: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 1, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, projectsList)
}

func (h *Handler) deleteProjectByID(ctx *gin.Context) {
	idParam := ctx.Query("id")
	if idParam == "" {
		log.Printf("id parameter is empty\n")
		ctx.JSON(http.StatusBadRequest, api.Error{Code: 1, Message: "empty id param"})
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id == 0 {
		log.Printf("can not get id: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, api.Error{Code: 1, Message: "invalid id param"})
		return
	}

	err = h.srvs.DeleteProjectByID(ctx, int64(id))
	if err != nil {
		log.Printf("can not delete Project where id = %d: %w", id, err)
		switch err {
		case custom_error.ErrProjectNotFound:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 3, Message: err.Error()})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 4, Message: err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (h *Handler) updateProjectByID(ctx *gin.Context) {
	idParam := ctx.Query("id")
	if idParam == "" {
		log.Printf("id parameter is empty\n")
		ctx.JSON(http.StatusBadRequest, api.Error{Code: 1, Message: "empty id param"})
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id == 0 {
		log.Printf("can not get id: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, api.Error{Code: 1, Message: "invalid id param"})
		return
	}

	var req entity.Projects
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 2, Message: custom_error.ErrInvalidInputBody.Error()})
		return
	}

	project, err := h.srvs.UpdateProjectByID(ctx, int64(id), req.Name)
	if err != nil {
		log.Printf("can not update Project where id = %d: %w", id, err)
		switch err {
		case custom_error.ErrProjectNotFound:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 3, Message: err.Error()})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 4, Message: err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, project)
}
