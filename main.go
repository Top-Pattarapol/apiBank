package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

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

	response, err := http.Get("http://jsonplaceholder.typicode.com/users")
	defer response.Body.Close()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Users
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(len(responseObject))
	fmt.Println(responseObject)

}

// func callService(url string) {

// 	response, err := http.Get(url)
// 	defer response.Body.Close()

// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
