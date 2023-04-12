package main

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	count      int    = 0
	firstHost  string = "http://127.0.0.1:8080"
	secondHost string = "http://127.0.0.1:8081"
)

func main() {
	rout := chi.NewRouter()
	rout.Post("/create", createUser)
	rout.Post("/make_friends", makeFriends)
	rout.Delete("/user", userDelete)
	rout.Get("/friends/{id}", friendsCheck)
	rout.Put("/{id}", updateAge)
	log.Fatalln("Ошибка прослушивания:", http.ListenAndServe(proxyAddr, rout))
}

func createUser(rw http.ResponseWriter, r *http.Request) {
	firstHost = "http://127.0.0.1:8080/create"
	secondHost = "http://127.0.0.1:8081/create"
	proxy(rw, r)
	return
}

func makeFriends(rw http.ResponseWriter, r *http.Request) {
	firstHost = "http://127.0.0.1:8080/make_friends"
	secondHost = "http://127.0.0.1:8081/make_friends"
	proxy(rw, r)
	return
}

func userDelete(rw http.ResponseWriter, r *http.Request) {
	firstHost = "http://127.0.0.1:8080/user"
	secondHost = "http://127.0.0.1:8081/user"
	proxy(rw, r)
	return
}

func friendsCheck(rw http.ResponseWriter, r *http.Request) {
	firstHost = "http://127.0.0.1:8080" + r.URL.Path
	secondHost = "http://127.0.0.1:8081" + r.URL.Path
	proxy(rw, r)
	return
}

func updateAge(rw http.ResponseWriter, r *http.Request) {
	firstHost = "http://127.0.0.1:8080" + r.URL.Path
	secondHost = "http://127.0.0.1:8081" + r.URL.Path
	proxy(rw, r)
	return
}

func proxy(rw http.ResponseWriter, r *http.Request) {
	host := firstHost
	if count == 1 {
		host = secondHost
	}

	newR, err := http.NewRequest(r.Method, host, r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}

	client := http.Client{}
	resp, err := client.Do(newR)
	defer resp.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write(respBody)

	if count == 1 {
		count--
	} else {
		count++
	}

	return
}
