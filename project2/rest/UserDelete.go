package rest

import (
	"awesomeProject/project1/database"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type UserIdDeleteStruct struct {
	TargetID string `json:"target_id"`
}

func UserDelete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	var del UserIdDeleteStruct

	if err = json.Unmarshal(content, &del); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	numDel, err := strconv.Atoi(del.TargetID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
	numMaxUserId, err := strconv.Atoi(database.MaxUserID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	// Проверка пользователя в мапе и удаление его ID у друзей
	if numMaxUserId >= numDel {
		for _, friendId := range database.MapUsers[del.TargetID].Friends {
			friend := database.MapUsers[friendId]
			friend.Friends = RemoveIDUser(friend.Friends, del.TargetID)
			database.MapUsers[friendId] = friend
		}
		rw.Write([]byte("Удален пользователь с именем: " + database.MapUsers[del.TargetID].Name + "\n"))
		delete(database.MapUsers, del.TargetID)
	} else {
		rw.Write([]byte("Пользователь с данным ID не найден\n"))
		rw.WriteHeader(http.StatusNotFound)
	}
	database.InFile(database.MapUsers)
	return
}

// RemoveIDUser удаляет ID удаленного юзера из массива Friends
func RemoveIDUser(arr []string, check string) []string {
	for i, v := range arr {
		if v == check {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
