/*
Docker Container for Go
https://www.youtube.com/watch?v=USbPCBi_d4U
https://www.youtube.com/watch?v=hjOXKmgilTo
*/
package main

import (
	"fmt"
	"net/http"
)

func main(){
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to the Go Web API!")
	fmt.Fprintf(response, "Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit: HomePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Ray Net"

	fmt.Fprintf(response, "A little bit about Ray Net...")
	fmt.Println("Endpoint Hit:", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.ListenAndServe(":8080", nil)
	
}