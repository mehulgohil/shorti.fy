package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/interfaces"
	"go.uber.org/zap"
)

type ShortifyReaderController struct {
	interfaces.IShortifyReaderService
	Logger *zap.Logger
}

// @Summary		Reader
// @Description	redirect to original url
// @Tags			shortify
// @Param			hashKey	path	string	true	"short url key"
// @Success		301
// @Failure		500
// @Router			/{hashKey} [get]
func (controller *ShortifyReaderController) ReaderController(ctx iris.Context) {
	params := ctx.Params()
	hashKey := params.Get("hashKey")

	originalURL, err := controller.Reader(hashKey)
	if err != nil {
		controller.Logger.Error(
			"READER: Error from service..",
			zap.String("hashKey", hashKey),
			zap.Error(err),
		)

		_ = ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(originalURL, iris.StatusMovedPermanently)
}
