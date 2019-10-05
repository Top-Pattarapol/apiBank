package main

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"net/http"
)

type Users []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Albums []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type CallService struct {
	*http.Response
	error
}

func main() {
	var users = &Users{}
	service := CallService{}
	err := service.Get("http://jsonplaceholder.typicode.com/users", users)
	fmt.Print(*users)
	fmt.Println(err)

}

func (e *CallService) Get(url string, model interface{}) error {
	e.Response, e.error = http.Get(url)
	if e.error != nil {
		return fmt.Errorf("network error : %w", e.error)
	}
	if e.Response.StatusCode != http.StatusOK {
		return fmt.Errorf("request error : %w", e.Response.StatusCode)
	}
	return nil
}

func (e *CallService) MapModel(model interface{}) error {
	data := json.NewDecoder(e.Response.Body)
	e.error = data.Decode(model)
	if e.error != nil {
		return fmt.Errorf("context wrapping: %w", e.error)
	}
	e.error = e.Response.Body.Close()
	if e.error != nil {
		return fmt.Errorf("close network error: %w", e.error)
	}
	return nil
}