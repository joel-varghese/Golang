package main

import(
	"fmt"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}