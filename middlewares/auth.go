package middlewares

import (
	"net/http"

	"github.com/agusheryanto182/go-todo/helpers"
	"github.com/agusheryanto182/go-todo/models/config"
)

type AuthMiddleware struct {
	Handler http.Handler
	Key     string
}

func NewAuthMiddleware(handler http.Handler, c *config.Global) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
		Key:     c.Api.Key,
	}
}

func (self *AuthMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if self.Key == req.Header.Get("x-api-key") {
		self.Handler.ServeHTTP(res, req)
	} else {
		helpers.WriteToResponseBodyError(res, http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
