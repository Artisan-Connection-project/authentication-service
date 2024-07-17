package handlers

import (
	"authentication-service/services"

	"github.com/sirupsen/logrus"
)

type MainHandler interface {
	Authentication() AuthenticationHandler
	User() UserHandler
	Token() TokenHandler
}

type mainHandler struct {
	authentication AuthenticationHandler
	user           UserHandler
	token          TokenHandler
}

func NewMainHandler(authSer services.AuthenticationService, userSer services.UserManagementService, tokenSer services.TokenService, log *logrus.Logger) MainHandler {
	return &mainHandler{
		authentication: NewAuthenticationHandler(authSer, log),
		user:           NewUserHandler(userSer, log),
		token:          NewTokenHandler(tokenSer, log),
	}
}

func (h *mainHandler) Authentication() AuthenticationHandler {
	return h.authentication
}

func (h *mainHandler) User() UserHandler {
	return h.user
}

func (h *mainHandler) Token() TokenHandler {
	return h.token
}
