package controller

import (
	"net/http"
	"strconv"

	"github.com/fondazionecrocereale/navegayaapi/data/request"
	"github.com/fondazionecrocereale/navegayaapi/data/response"
	"github.com/fondazionecrocereale/navegayaapi/helper"
	"github.com/fondazionecrocereale/navegayaapi/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ClipsController struct {
	clipsService service.ClipsService
}

func NewClipsController(service service.ClipsService) *ClipsController {
	return &ClipsController{
		clipsService: service,
	}
}

// CreateClips		godoc
// @Summary			Create clips
// @Description		Save clips data in Db.
// @Param			clips body request.CreateClipsRequest true "Create clips"
// @Produce			application/json
// @Clips			clips
// @Success			200 {object} response.Response{}
// @Router			/clips [post]
func (controller *ClipsController) Create(ctx *gin.Context) {
	log.Info().Msg("create clips")
	createClipsRequest := request.CreateClipsRequest{}
	err := ctx.ShouldBindJSON(&createClipsRequest)
	helper.ErrorPanic(err)

	controller.clipsService.Create(createClipsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateClips		godoc
// @Summary			Update clips
// @Description		Update clips data.
// @Param			clipId path string true "update clips by id"
// @Param			clips body request.CreateClipsRequest true  "Update clips"
// @Clips			clips
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/clips/{clipId} [patch]
func (controller *ClipsController) Update(ctx *gin.Context) {
	log.Info().Msg("update clips")
	updateClipsRequest := request.UpdateClipsRequest{}
	err := ctx.ShouldBindJSON(&updateClipsRequest)
	helper.ErrorPanic(err)

	clipId := ctx.Param("clipId")
	id, err := strconv.Atoi(clipId)
	helper.ErrorPanic(err)
	updateClipsRequest.Id = id

	controller.clipsService.Update(updateClipsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteClips		godoc
// @Summary			Delete clips
// @Description		Remove clips data by id.
// @Produce			application/json
// @Clips			clips
// @Success			200 {object} response.Response{}
// @Router			/clips/{clipID} [delete]
func (controller *ClipsController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete clips")
	clipId := ctx.Param("clipId")
	id, err := strconv.Atoi(clipId)
	helper.ErrorPanic(err)
	controller.clipsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdClips 		godoc
// @Summary				Get Single clips by id.
// @Param				clipId path string true "update clips by id"
// @Description			Return the tahs whoes clipId valu mathes id.
// @Produce				application/json
// @Clips				clips
// @Success				200 {object} response.Response{}
// @Router				/clips/{clipId} [get]
func (controller *ClipsController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid clips")
	clipId := ctx.Param("clipId")
	id, err := strconv.Atoi(clipId)
	helper.ErrorPanic(err)

	clipResponse := controller.clipsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   clipResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllClips 		godoc
// @Summary			Get All clips.
// @Description		Return list of clips.
// @Clips			clips
// @Success			200 {obejct} response.Response{}
// @Router			/clips [get]
func (controller *ClipsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll clips")
	clipResponse := controller.clipsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   clipResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
