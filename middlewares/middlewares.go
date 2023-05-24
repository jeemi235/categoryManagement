package middlewares

import (
	e "categorymanagement/api/utils/errrors"
	"categorymanagement/database"
	"context"
	"encoding/json"
	"net/http"
)

// With this function we are creating context of DB(database)
func DbContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.Connect()
		ctx := context.WithValue(r.Context(), "database", db)
		defer db.Close()
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// it will encode the data of output
func ResponseWithJsonPayload(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}

// This function Authorize the user and then call the api
func Authuser(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("id")
		userid := "863611762007769089"
		if id != userid {
			e.ErrorGenerator(w, e.Unauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
