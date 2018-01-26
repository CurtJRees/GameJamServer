package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"encoding/json"
	"github.com/CurtJRees/gamejamserver/model"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", NotImplemented)
	router.POST("/command/", SendCommand)

	return router
}

var NotImplemented = httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Not Implemented"))
})

var SendCommand = httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	var comm model.Command

	err := decoder.Decode(&comm)
	if err != nil {
		panic(err)
	}

	log.Println("Recieved command - " + model.CommandToString(comm))
})