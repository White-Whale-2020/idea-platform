/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:		source file for main module
	Language:			Go

	Authors: 			David Cho
	 							Mike Roper
	 							Kevon Nelson
	 							Kyle Thomas

	Organization:	White-Whale-2020 (Coding Cougs WSUTC)
	Event:				Crimson Code 2020
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	log.Println("Welcome to the IDP (IDea Platform)")
	log.Println("Authors: David Cho, Mike Groper, Kevon Nelson, Kyle Thomas")
	log.Println("Version: beta")

	port, exists := os.LookupEnv("GOPORT")
	if !exists {
		port = defaultPort
	}

	http.HandleFunc("/", frontend)
	http.HandleFunc("/api/", api)

	/*
		// this will serve css and js files later
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	*/

	log.Println("Web server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
		notFound(w)
	}
}

func apiDocs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("{\"api-version\":\"1.0\",\"endpoint\":\"api\",\"endpoint\":[\"docs\",\"going\",\"hackathon\",\"idea\",\"member\",\"room\",\"team\",\"vehicle\"]}"))
	case "POST": // docs shouldn't have a post method but keeping for the "printStuff()"
		printStuff(r)
	default:
		notFound(w)
	}
}

func apiGoing(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(GoingGet("sampleJSON/going.json"))
	case "POST":
		w.Write(GoingPut(r))
	default:
		notFound(w)
	}
}

func apiHackathon(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(HackathonGet("sampleJSON/hackathon.json"))
	case "POST":
		w.Write(HackathonPut(r))
	default:
		notFound(w)
	}
}

func apiIdea(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(IdeaGet("sampleJSON/idea.json"))
	case "POST":
		w.Write(IdeaPut(r))
	default:
		notFound(w)
	}
}

func apiMember(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(MemberGet("sampleJSON/member.json"))
	case "POST":
		w.Write(MemberPut(r))
	default:
		notFound(w)
	}
}

func apiRoom(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(RoomGet("sampleJSON/room.json"))
	case "POST":
		w.Write(RoomPut(r))
	default:
		notFound(w)
	}
}

func apiTeam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(TeamGet("sampleJSON/team.json"))
	case "POST":
		w.Write(TeamPut(r))
	default:
		notFound(w)
	}
}

func apiVehicle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(VehicleGet("sampleJSON/vehicle.json"))
	case "POST":
		w.Write(VehiclePut(r))
	default:
		notFound(w)
	}
}

func notFound(w http.ResponseWriter) {
	w.Write([]byte("{\"status\":\"404 not found\"}"))
}

func printStuff(r *http.Request) {
	fmt.Println("<<<<< Header >>>>>>>")
	for key, element := range r.Header {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
	fmt.Println("<<<<< Body >>>>>>>")
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

}
