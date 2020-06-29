package controllers

import (
	"github.com/thaibuixuanDEV/forum/api/responses"
	"net/http"
)

func (server *Server) Home(writer http.ResponseWriter, request *http.Request) {
	responses.JSON(writer, http.StatusOK, "Welcome to my API")
}

