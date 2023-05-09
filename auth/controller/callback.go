package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/auth/authenticator"
	"net/http"
)

type CallbackHandler struct {
	Auth *authenticator.Authenticator
}

func (c *CallbackHandler) Callback(ctx iris.Context) {

	if ctx.URLParam("state") != state {
		ctx.StopWithJSON(http.StatusBadRequest, "Invalid state parameter.")
		return
	}

	// Exchange an authorization code for a token.
	token, err := c.Auth.Exchange(ctx.Request().Context(), ctx.URLParam("code"))
	if err != nil {
		ctx.StopWithJSON(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
		return
	}

	idToken, err := c.Auth.VerifyIDToken(ctx.Request().Context(), token)
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	TOKEN = token.AccessToken
	PROFILE = profile["email"].(string)

	ctx.SetCookieKV("logged_id_email", profile["email"].(string), iris.CookieHTTPOnly(false))

	//// Redirect to logged in page.
	//ctx.Redirect("/user", http.StatusTemporaryRedirect)
}
