package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}