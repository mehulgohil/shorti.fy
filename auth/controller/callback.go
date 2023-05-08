package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/shorti.fy/auth/authenticator"
	"net/http"
)

type CallbackHandler struct {
	Auth *authenticator.Authenticator
}

func (c *CallbackHandler) Callback(ctx iris.Context) {
	session := sessions.Get(ctx)

	if ctx.URLParam("state") != session.Get("state") {
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
	//fmt.Println(token)
	fmt.Println(profile)
	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	fmt.Println("============")
	fmt.Println(session.Get("profile"))
	fmt.Println("============")
	// Redirect to logged in page.
	ctx.Redirect("/user", http.StatusTemporaryRedirect)
}
