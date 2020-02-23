/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:		source file for main module
	Language:			Go

	Authors: 			David Cho
	 							Mike Groper
	 							Kevon Nelson
	 							Kyle Thomas

	Organization:	White-Whale-2020 (Coding Cougs WSUTC)
	Event:				Crimson Code 2020
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

	/*
		// this will serve css and js files later
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
	fmt.Fprint(w, "Front end/root url!\n")
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	switch r.URL.Path {
	case "/api/", "/api/docs":
		apiDocs(w, r)
	case "/api/going":
		apiGoing(w, r)
	case "/api/hackathon":
		apiHackathon(w, r)
	case "/api/idea":
		apiIdea(w, r)
	case "/api/member":
		apiMember(w, r)
	case "/api/room":
		apiRoom(w, r)
	case "/api/team":
		apiTeam(w, r)
	case "/api/vehicle":
		apiVehicle(w, r)
	default:
		notFound(w, r)
	}
}

func apiDocs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("{\"api-version\":\"1.0\",\"endpoint\":\"api\",\"endpoint\":[\"docs\",\"going\",\"hackathon\",\"idea\",\"member\",\"room\",\"team\",\"vehicle\"]}"))
	case "POST":
	default:
		notFound(w, r)
	}
}

func apiGoing(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestGoingJSON("sampleJSON/going.json"))
	case "POST":
		fmt.Fprint(w, "Ideas Post!\n")
	default:
		notFound(w, r)
	}
}

func apiHackathon(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestHackathonJSON("sampleJSON/hackathon.json"))
	case "POST":
		fmt.Fprint(w, "Ideas Post!\n")
	default:
		notFound(w, r)
	}
}

func apiIdea(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestIdeaJSON("sampleJSON/idea.json"))
	case "POST":
		fmt.Fprint(w, "Ideas Post!\n")
	default:
		notFound(w, r)
	}
}

func apiMember(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestMemberJSON("sampleJSON/member.json"))
	case "POST":
		fmt.Fprint(w, "Users Post!\n")
	default:
		notFound(w, r)
	}
}

func apiRoom(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestRoomJSON("sampleJSON/room.json"))
	case "POST":
		fmt.Fprint(w, "Ideas Post!\n")
	default:
		notFound(w, r)
	}
}

func apiTeam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestTeamJSON("sampleJSON/team.json"))
	case "POST":
		fmt.Fprint(w, "Teams Post!\n")
	default:
		notFound(w, r)
	}
}

func apiVehicle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TestVehicleJSON("sampleJSON/vehicle.json"))
	case "POST":
		fmt.Fprint(w, "Users Post!\n")
	default:
		notFound(w, r)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"status\":\"404 not found\"}"))
}
