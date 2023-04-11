package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/interfaces"
	"github.com/mehulgohil/shorti.fy/models"
)

type ShortifyController struct {
	interfaces.IShortifyService
}

func (controller *ShortifyController) ReaderController(ctx iris.Context) {
	params := ctx.Params()
	hashedValue := params.Get("hashedValue")

	originalURL, err := controller.Reader(hashedValue)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(originalURL, iris.StatusMovedPermanently)
}

func (controller *ShortifyController) WriterController(ctx iris.Context) {
	var requestBody models.WriterRequest
	err := ctx.ReadJSON(&requestBody)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
	}

	newShortURL, err := controller.Writer(requestBody.LongURL)
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
