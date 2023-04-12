package rest

import (
	"awesomeProject/project1/database"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func FriendsCheck(rw http.ResponseWriter, r *http.Request) {
	userIdFriendsCheck := chi.URLParam(r, "id")

	if userIdFriendsCheck <= database.MaxUserID {
		friends := ""
		for _, v := range database.MapUsers[userIdFriendsCheck].Friends {
			friends += database.MapUsers[v].Name + " "
		}
		rw.Write([]byte("Друзья пользователя: " + friends + "\n"))
	} else {
		rw.Write([]byte("Пользователь не найден\n"))
		rw.WriteHeader(http.StatusNotFound)
	}
	return
}
