package main

import (
	rest2 "awesomeProject/project1/rest"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	fmt.Println("30.5 Практическая работа")
	fmt.Println("-------------------")

	rout := chi.NewRouter()
	rout.Post("/create", rest2.CreateUser)
	rout.Post("/make_friends", rest2.MakeFriends)
	rout.Delete("/user", rest2.UserDelete)
	rout.Get("/friends/{id}", rest2.FriendsCheck)
	rout.Put("/{id}", rest2.UpdateAge)

	err := http.ListenAndServe("localhost:8081", rout)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
