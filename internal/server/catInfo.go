package server

import (
	"catalk/instructions"
	"catalk/utils"
	"net/http"
)

func (s *Server) CatBreeds(w http.ResponseWriter, r *http.Request) {
	catBreedsSlice := make([]string, 0)
	for _, breed := range instructions.CatBreedsMap {
		catBreedsSlice = append(catBreedsSlice, breed)
	}

	utils.SuccessResponse(w, http.StatusOK, catBreedsSlice)
}
