package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"encoding/json"
)


func main(){
	fmt.Println("welcome to learn")
	PerformPutJsonRequest()
}

func PerformPutJsonRequest(){
	const url = "https://reqres.in/api/users/2"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"name": "AAAAAMMMMOOOLLLLL",
			"job": "Learning new"
		}
	`)
	response, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
