/*
 * Application: 	IDP (IDea Platform)
 * File:					idp.go
 * Description:		source file for main module
 * Language:			Go
 *
 * Authors: 			David Cho
 *								Mike Groper
 *								Kevon Nelson
 *								Kyle Thomas
 *
 * Organization:	White-Whale-2020 (Coding Cougs WSUTC)
 * Event:					Crimson Code 2020
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	log.Println("Welcome to the IDP (IDea Platform)")
	log.Println("Authors: David Cho, Mike Groper, Kevon Nelson, Kyle Thomas")
	log.Println("Version: beta")

	http.HandleFunc("/", handler)
	//http.HandleFunc("/api", )
	//http.HandleFunc("/api", )
	log.Println("Web server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
