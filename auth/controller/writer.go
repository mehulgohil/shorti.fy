package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
)

type WriterHandler struct{}

func (w *WriterHandler) WriterRedirect(ctx iris.Context) {
	client := &http.Client{}
	req, err := http.NewRequest(ctx.Request().Method, "http://localhost:3000/v1/shorten", ctx.Request().Body)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}
	fmt.Println(TOKEN)
	req.Header.Add("Authorization", "Bearer "+TOKEN)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}

	ctx.StopWithJSON(res.StatusCode, respBody)
}
