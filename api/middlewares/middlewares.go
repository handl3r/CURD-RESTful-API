package middlewares

import (
	"errors"
	"github.com/thaibuixuanDEV/forum/api/auth"
	"github.com/thaibuixuanDEV/forum/api/responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next(writer, request)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := auth.TokenValid(request)
		if err != nil {
			responses.ERROR(writer, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(writer, request)
	}
}
