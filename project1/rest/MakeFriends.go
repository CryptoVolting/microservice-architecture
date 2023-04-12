package rest

import (
	"awesomeProject/project1/database"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type AddFriendStruct struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

func MakeFriends(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	var f AddFriendStruct

	if err = json.Unmarshal(content, &f); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	numSourceId, err := strconv.Atoi(f.SourceId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
	numTargetId, err := strconv.Atoi(f.TargetId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
	numMaxUserId, err := strconv.Atoi(database.MaxUserID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
	// Проверка пользователей в мапе и добавление ID друзей в User.Friends
	if numMaxUserId >= numSourceId && numMaxUserId >= numTargetId {
		sourFr := database.MapUsers[f.SourceId]
		targFr := database.MapUsers[f.TargetId]

		sourFr.Friends = append(sourFr.Friends, f.TargetId)
		targFr.Friends = append(targFr.Friends, f.SourceId)

		database.MapUsers[f.SourceId] = sourFr
		database.MapUsers[f.TargetId] = targFr

		get := database.MapUsers[f.SourceId].Name + " и " + database.MapUsers[f.TargetId].Name + " теперь друзья!\n"
		rw.Write([]byte(get))
	} else {
		rw.Write([]byte("Пользователи с данными ID не найдены\n"))
		rw.WriteHeader(http.StatusNotFound)
	}
	database.InFile(database.MapUsers)
	return
}
