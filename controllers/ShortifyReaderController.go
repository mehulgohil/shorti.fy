package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/interfaces"
)

type ShortifyReaderController struct {
	interfaces.IShortifyReaderService
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
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(originalURL, iris.StatusMovedPermanently)
}
