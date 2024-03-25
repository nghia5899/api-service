package router

import (
	"apiService/service"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/view", service.View).Methods("POST")
	router.HandleFunc("/update/cookie", service.Upload).Methods("POST")
	router.HandleFunc("/update/proxy", service.UploadProxy).Methods("POST")
	return router
}
