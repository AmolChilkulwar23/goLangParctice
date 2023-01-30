package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main(){
	fmt.Println("welcome to learn")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest(){
	const url = "https://reqres.in/api/users"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"name": "AAAAAMMMMOOOLLLLL",
			"job": "Learning"
		}
	`)
	response, err := http.Post(url,"application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
