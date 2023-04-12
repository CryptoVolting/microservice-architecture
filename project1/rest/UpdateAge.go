package rest

import (
	"awesomeProject/project1/database"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type UpdateAgeStruct struct {
	NewAge string `json:"new_age"`
}

func UpdateAge(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id")
	if userId <= database.MaxUserID {
		content, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}

		var newAge UpdateAgeStruct

		if err = json.Unmarshal(content, &newAge); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		userUpdate := database.MapUsers[userId]
		userUpdate.Age = newAge.NewAge
		database.MapUsers[userId] = userUpdate
		rw.Write([]byte("Возраст пользователя успешно обновлен.\n"))
	} else {
		rw.Write([]byte("Пользователь не найден\n"))
		rw.WriteHeader(http.StatusNotFound)
	}
	database.InFile(database.MapUsers)
	return
}
