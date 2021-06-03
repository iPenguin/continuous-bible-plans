/**
 *
 **/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"

	"github.com/ipenguin/continuous-bible-plans/internal/account"
	"github.com/ipenguin/continuous-bible-plans/internal/dashboard"
	"github.com/ipenguin/continuous-bible-plans/internal/index"
	"github.com/ipenguin/continuous-bible-plans/internal/prayer"
	"github.com/ipenguin/continuous-bible-plans/internal/read"
	"github.com/ipenguin/continuous-bible-plans/internal/resources"
	"github.com/ipenguin/continuous-bible-plans/internal/signup"
	"github.com/ipenguin/continuous-bible-plans/internal/tags"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func main() {

	router := mux.NewRouter()

	// static resource files
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

	for _, page := range signup.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, signup.GeneratePage)
	}

	for _, page := range account.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, account.GeneratePage)
	}

	for _, page := range read.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, read.GeneratePage)
	}

	for _, page := range dashboard.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, dashboard.GeneratePage)
	}

	for _, page := range tags.PageNames() {
		log.Println("Setting index route: " + page)
		router.HandleFunc(page, tags.GeneratePage)
	}

	http.ListenAndServe(":8080", router)
}
