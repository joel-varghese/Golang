package main

import(
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello there, WTF %s!",r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}