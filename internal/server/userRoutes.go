package server

import (
	"catalk/internal/users"
	"net/http"
)

func (s *Server) AddUser(w http.ResponseWriter, r *http.Request) {
	users.InsertUser(&users.UserEntity{
		Email:      "test2@gmail.com",
		Username:   "testJa",
		PictureURL: "test.img",
		ProviderID: 1,
	})
}
