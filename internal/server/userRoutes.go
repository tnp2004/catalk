package server

import (
	"catalk/internal/users"
	"catalk/utils"
	"net/http"
)

func (s *Server) AddUser(w http.ResponseWriter, r *http.Request) {
	reqBody := new(users.NewUserModel)
	if err := utils.ReadReqBody(r, reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	user := users.NewUser(s.config.Database)
	if _, err := user.InsertUser(reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}
