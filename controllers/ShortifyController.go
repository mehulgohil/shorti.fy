package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/interfaces"
)

type ShortifyController struct {
	interfaces.IShortifyService
}

func (controller *ShortifyController) ReaderController(ctx iris.Context) {
	originalURL, err := controller.Reader("https://shorturl")
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"error": err.Error(),
		})
		return
	}

	_ = ctx.JSON(iris.Map{
		"redirectURL": originalURL,
	})
}

func (controller *ShortifyController) WriterController(ctx iris.Context) {
	newShortURL, err := controller.Writer("https://longurl")
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"error": err.Error(),
		})
		return
	}

	_ = ctx.JSON(iris.Map{
		"newURL": newShortURL,
	})
}
