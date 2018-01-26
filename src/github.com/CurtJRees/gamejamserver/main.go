package main

import(
	"github.com/CurtJRees/gamejamserver/router"
	"log"
	"net/http"
)

func main() {

	r := router.NewRouter()

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}