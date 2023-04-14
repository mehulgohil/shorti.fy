package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/interfaces"
	"github.com/mehulgohil/shorti.fy/models"
)

type ShortifyWriterController struct {
	interfaces.IShortifyWriterService
}

// @Summary		Writer
// @Description	shorten a long url
// @Tags			shortify
// @Accept			json
// @Produce		json
// @Param			data	body		models.WriterRequest	true	"writer request body"
// @Success		200		{object}	models.WriterResponse
// @Failure		400
// @Failure		500
// @Router			/shorten [post]
func (controller *ShortifyWriterController) WriterController(ctx iris.Context) {
	var requestBody models.WriterRequest
	err := ctx.ReadJSON(&requestBody)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
	}

	newShortURL, err := controller.Writer(requestBody.LongURL, requestBody.UserEmail)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"error": err.Error(),
		})
		return
	}

	_ = ctx.JSON(models.WriterResponse{
		LongURL:  requestBody.LongURL,
		ShortURL: newShortURL,
	})
}
