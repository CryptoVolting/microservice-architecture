package rest

import (
	"awesomeProject/project1/database"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	var u database.User

	if err = json.Unmarshal(content, &u); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	database.MapUsers[database.MaxUserID] = u
	get := "UserID " + u.Name + ":" + database.MaxUserID + "\n"
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(get))

	database.MaxUserID = userIDPlusOne(database.MaxUserID)
	database.InFile(database.MapUsers)
	return
}

// userIDPlusOne увеличивает значение MaxUserID на единицу
func userIDPlusOne(userID string) string {
	id, err := strconv.Atoi(userID)
	if err != nil {
		panic(err.Error())
	}
	id++
	return strconv.Itoa(id)
}
