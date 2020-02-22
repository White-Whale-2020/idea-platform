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

	http.HandleFunc("/", frontend)
	http.HandleFunc("/api/", api)

	/* this will serve css and js files later
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	*/

	log.Println("Web server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func frontend(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Front end/root url!")
}

func api(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/", "/api/docs":
		fmt.Fprint(w, "API endpoint!\n")
		fmt.Fprint(w, "Eventually put api docs here?")
	case "/api/teams":
		fmt.Fprint(w, "Teams!")
	case "/api/users":
		fmt.Fprint(w, "Users!")
	case "/api/ideas":
		fmt.Fprint(w, "ideas!")
	default:
		http.NotFound(w, r)
	}
}
