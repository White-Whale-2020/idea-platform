/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "team" api endpoint
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TeamsS struct {
	Endpoint   string  `json:"endpoint"`
	APIversion string  `json:"api-version"`
	Team       []TeamS `json:"teams"`
}

type TeamS struct {
	Hackathon string   `json:"hackathon"`
	Idea      string   `json:"idea"`
	Members   []string `json:"members"`
	Awards    string   `json:"awards"`
}

func TeamGet(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var ts TeamsS
	json.Unmarshal(rtn, &ts)
	rtn, err = json.Marshal(ts)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}

func TeamPut(r *http.Request) []byte {
	var tmpMbr TeamS
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading json")
		log.Print(err)
		fmt.Println()
		return []byte("{\"status\":\"404 not found\"}")
	}
	json.Unmarshal(body, &tmpMbr)
	// put the json into a database now
	return []byte("{\"status\":\"200 ok\"}")
}
