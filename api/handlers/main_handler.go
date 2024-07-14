package handlers

import "authentication-service/services"

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

func NewMainHandler(authSer services.AuthenticationService, userSer services.UserManagementService, tokenSer services.TokenService) MainHandler {
	return &mainHandler{
		authentication: NewAuthenticationHandler(authSer),
		user:           NewUserHandler(userSer),
		token:          NewTokenHandler(tokenSer),
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
