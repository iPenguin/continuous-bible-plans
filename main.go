/**
 *
 **/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"

	"github.com/ipenguin/continuous-bible-plans/internal/index"
	"github.com/ipenguin/continuous-bible-plans/internal/prayer"
	"github.com/ipenguin/continuous-bible-plans/internal/resources"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func main() {

	router := mux.NewRouter()

	//TODO: get from config
	//router.Host("")
	for _, page := range resources.PageNames() {
		log.Println("Setting resource routes: " + page)
		router.PathPrefix(page + "/").Handler(http.FileServer(http.Dir("./assets/")))
	}

	for _, page := range index.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, index.GeneratePage)
	}

	for _, page := range prayer.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, prayer.GeneratePage)
	}

	http.ListenAndServe(":8080", router)
}
