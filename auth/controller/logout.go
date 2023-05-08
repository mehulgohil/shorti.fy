package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/shorti.fy/auth/config"
	"net/http"
	"net/url"
)

type LogoutHandler struct{}

func (l *LogoutHandler) Logout(ctx iris.Context) {
	session := sessions.Get(ctx)

	session.Destroy()
	logoutUrl, err := url.Parse("https://" + config.EnvVariables.Auth0Domain + "/v2/logout")
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	scheme := "http"
	if ctx.Request().TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + ctx.Request().Host)
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", config.EnvVariables.Auth0ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}
