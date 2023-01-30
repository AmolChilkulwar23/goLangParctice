package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
	"encoding/json"
)


func main(){
	fmt.Println("welcome to learn")
	PerformPutJsonRequest()
}

func PerformPutJsonRequest(){
	const url = "https://reqres.in/api/users/2"

	//fake json payload
	payload, err := json.Marshal(map[string]interface{} {
		"name": "AAAAAMMMMOOOLLLLL",
		"job": "zion resident",
	})
	response, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
