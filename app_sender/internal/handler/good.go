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

func (h *Handler) createGood(ctx *gin.Context) {
	projectId, _ := strconv.ParseInt(ctx.DefaultQuery("projectId", "0"), 10, 64)
	if projectId == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 1, Message: custom_error.ErrInvalidURLHead.Error()})
		return
	}

	var req entity.Goods
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 1, Message: custom_error.ErrInvalidInputBody.Error()})
		return
	}

	req.ProjectID = projectId

	goods, err := h.srvs.CreateGood(ctx, &req)
	if err != nil {
		log.Printf("can not create project: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 2, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, goods)
}

func (h *Handler) getAllGoods(ctx *gin.Context) {
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.Error{Code: 1, Message: "invalid id param"})
		return
	}

	goodsList, err := h.srvs.GetAllGoods(ctx, limit, offset)
	if err != nil {
		log.Printf("can not get all projects: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 1, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, goodsList)
}

func (h *Handler) deleteGoodByID(ctx *gin.Context) {
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

	err = h.srvs.DeleteGoodByID(ctx, id)
	if err != nil {
		log.Printf("can not delete Project where id = %d: %v", id, err)
		switch err {
		case custom_error.ErrGoodNotFound:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 3, Message: err.Error()})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Code: 4, Message: err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, api.RemoveGoods{Id: id, Removed: true})
}

func (h *Handler) updateGoodByID(ctx *gin.Context) {
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

	var req entity.Goods
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 2, Message: custom_error.ErrInvalidInputBody.Error()})
		return
	}

	req.ID = id

	goods, err := h.srvs.UpdateGoodByID(ctx, id, &req)
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

	ctx.JSON(http.StatusOK, goods)
}

func (h *Handler) reprioritiize(ctx *gin.Context) {
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

	var req api.PayloadNewPriority
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Code: 2, Message: custom_error.ErrInvalidInputBody.Error()})
		return
	}

	priorities, err := h.srvs.Reprioritiize(ctx, id, req.NewPriority)
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

	ctx.JSON(http.StatusOK, priorities)
}
