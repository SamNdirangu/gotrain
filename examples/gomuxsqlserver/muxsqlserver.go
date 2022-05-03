package gomuxsqlserver

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StartServer(port int) {
	router := mux.NewRouter()
	http.Handle("/", router)

	var handle string = ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(handle, router))
}
