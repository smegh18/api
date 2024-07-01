package handlers

import "book_seller/domain"

type UserHandler struct {
	UserDomain domain.Service
}

func NewUserHandler(userDomain domain.Service) *UserHandler {
	return &UserHandler{UserDomain: userDomain}
}
