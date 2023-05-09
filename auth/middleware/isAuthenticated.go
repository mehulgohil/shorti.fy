package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"net/http"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated(ctx iris.Context) {
	userCookie := ctx.GetCookie("logged_id_email")
	fmt.Println(userCookie)
	session := sessions.Get(ctx)
	profileToken := session.Get("profile")
	if profileToken == nil {
		ctx.Redirect("/login", http.StatusSeeOther)
	} else {
		ctx.Next()
	}
}
