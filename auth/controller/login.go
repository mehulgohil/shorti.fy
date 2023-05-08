package controller

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/shorti.fy/auth/authenticator"
	"net/http"
)

type LoginHandler struct {
	Auth *authenticator.Authenticator
}

func (l *LoginHandler) Login(ctx iris.Context) {
	state, err := generateRandomState()
	if err != nil {
		ctx.StopWithPlainError(http.StatusInternalServerError, err)
		return
	}

	// Save the state inside the session.
	session := sessions.Get(ctx)
	session.Set("state", state)

	ctx.Redirect(l.Auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
