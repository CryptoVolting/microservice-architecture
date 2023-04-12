package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	log.Println(createUserTest(t))
	log.Println(makeFriendsTest(t))
	log.Println(userDeleteTest(t))
	log.Println(friendsCheckTest(t))
	log.Println(ageUpdateTest(t))
}

// countCheck проверяет значение count на работу прокси.
func countCheck(t *testing.T, i int) string {
	if i%2 == 0 {
		if count == 0 {
			t.Log("count не изменил свое значение")
			t.Fail()
			return "Fail!"
		}
	} else {
		if count == 1 {
			t.Log("count не изменил свое значение")
			t.Fail()
			return "Fail!"
		}
	}
	return "Pass!"
}

func createUserTest(t *testing.T) string {
	srv := httptest.NewServer(http.HandlerFunc(createUser))
	outputTest := "Fail!"
	for i := 0; i < 6; i++ {
		body := strings.NewReader(`{"name":"Masha","age":"26","friends":[]}`)
		contentType := "application/json"
		urlCreate := srv.URL + "/create"

		resp, err := http.Post(urlCreate, contentType, body)
		if err != nil {
			t.Log("Ошибка отправки Post запроса:", err.Error())
			t.Fail()
			return "createUserTest " + outputTest
		}

		respByte, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Log("Ошибка чтения resp.Body:", err.Error())
			t.Fail()
			return "createUserTest " + outputTest
		}

		if i == 0 {
			respTest := "UserID Masha:0\n"
			if string(respByte) != respTest {
				t.Log("Ответ сервера не соответствует запросу")
				t.Fail()
				return "createUserTest " + outputTest
			}
		}

		outputTest = countCheck(t, i)
	}
	return "createUserTest " + outputTest
}

func makeFriendsTest(t *testing.T) string {
	srv := httptest.NewServer(http.HandlerFunc(makeFriends))
	outputTest := "Fail!"
	for i := 0; i < 4; i++ {
		var body *strings.Reader
		if i < 2 {
			body = strings.NewReader(`{"source_id":"0","target_id":"1"}`)
		} else {
			body = strings.NewReader(`{"source_id":"0","target_id":"2"}`)
		}
		contentType := "application/json"
		urlTest := srv.URL + "/make_friends"

		resp, err := http.Post(urlTest, contentType, body)
		if err != nil {
			t.Log("Ошибка отправки Post запроса:", err.Error())
			t.Fail()
			return "makeFriendsTest " + outputTest
		}

		respByte, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Log("Ошибка чтения resp.Body:", err.Error())
			t.Fail()
			return "makeFriendsTest " + outputTest
		}

		respTest := "Masha и Masha теперь друзья!\n"
		if string(respByte) != respTest {
			t.Log("Ответ сервера не соответствует запросу")
			t.Fail()
			return "makeFriendsTest " + outputTest
		}

		outputTest = countCheck(t, i)
	}
	return "makeFriendsTest " + outputTest
}

func userDeleteTest(t *testing.T) string {
	srv := httptest.NewServer(http.HandlerFunc(userDelete))
	outputTest := "Fail!"
	for i := 0; i < 2; i++ {
		body := strings.NewReader(`{"target_id":"1"}`)
		urlTest := srv.URL + "/user"

		r, _ := http.NewRequest(http.MethodDelete, urlTest, body)
		client := srv.Client()
		resp, err := client.Do(r)
		if err != nil {
			t.Log("Ошибка получения ответа от сервера")
			t.Fail()
			return "userDeleteTest " + outputTest
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Log("Ошибка чтения ответа от сервера")
			t.Fail()
			return "userDeleteTest " + outputTest
		}

		respTest := "Удален пользователь с именем: Masha\n"
		if string(respBody) != respTest {
			t.Log("Ответ сервера не соответствует запросу")
			t.Fail()
			return "userDeleteTest " + outputTest
		}
		outputTest = countCheck(t, i)
	}
	return "userDeleteTest " + outputTest
}

func friendsCheckTest(t *testing.T) string {
	srv := httptest.NewServer(http.HandlerFunc(friendsCheck))
	outputTest := "Fail!"
	for i := 0; i < 2; i++ {
		urlTest := srv.URL + "/friends/2"

		resp, err := http.Get(urlTest)
		if err != nil {
			t.Log("Ошибка отправки Get запроса:", err.Error())
			t.Fail()
			return "friendsCheckTest " + outputTest
		}

		respByte, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Log("Ошибка чтения resp.Body:", err.Error())
			t.Fail()
			return "friendsCheckTest " + outputTest
		}

		respTest := "Друзья пользователя: Masha \n"
		if string(respByte) != respTest {
			t.Log("Ответ сервера не соответствует запросу")
			t.Fail()
			return "friendsCheckTest " + outputTest
		}

		outputTest = countCheck(t, i)
	}
	return "friendsCheckTest " + outputTest
}

func ageUpdateTest(t *testing.T) string {
	srv := httptest.NewServer(http.HandlerFunc(updateAge))
	outputTest := "Fail!"
	for i := 0; i < 2; i++ {
		body := strings.NewReader(`{"new_age":"28"}`)
		urlTest := srv.URL + "/0"

		r, _ := http.NewRequest(http.MethodPut, urlTest, body)
		client := srv.Client()
		resp, err := client.Do(r)
		if err != nil {
			t.Log("Ошибка получения ответа от сервера")
			t.Fail()
			return "ageUpdateTest " + outputTest
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Log("Ошибка чтения ответа от сервера")
			t.Fail()
			return "ageUpdateTest " + outputTest
		}

		respTest := "Возраст пользователя успешно обновлен.\n"
		if string(respBody) != respTest {
			t.Log("Ответ сервера не соответствует запросу")
			t.Fail()
			return "ageUpdateTest " + outputTest
		}
		outputTest = countCheck(t, i)
	}
	return "ageUpdateTest " + outputTest
}
