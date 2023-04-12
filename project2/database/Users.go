package database

import (
	"encoding/json"
	"log"
	"os"
)

var MaxUserID = "0"
var MapUsers = make(map[string]User)

type User struct {
	Name    string   `json:"name"`
	Age     string   `json:"age"`
	Friends []string `json:"friends"`
}

// InFile Создание или перезапись файла data.json
func InFile(mapUsers map[string]User) {

	rawDataOut, err := json.MarshalIndent(&mapUsers, "", "")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = os.WriteFile("data.json", rawDataOut, 0777)
	if err != nil {
		log.Fatal("Cannot write updated data file:", err)
	}
}
