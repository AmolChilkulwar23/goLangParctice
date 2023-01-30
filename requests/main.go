package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


const url = "https://loc.dev"

func main(){
	fmt.Println("web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response : %T\n", response)

	data, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
		panic(err)
	}
	datacontent := string(data)
	fmt.Println(datacontent)


}