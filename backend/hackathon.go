/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "hackathon" api endpoint
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HackathonsS struct {
	Endpoint   string       `json:"endpoint"`
	APIversion string       `json:"api-version"`
	Hackathon  []HackathonS `json:"hackathons"`
}

type HackathonS struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func HackathonGet(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var hs HackathonsS
	json.Unmarshal(rtn, &hs)
	rtn, err = json.Marshal(hs)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}

func HackathonPut(r *http.Request) []byte {
	var tmpMbr HackathonS
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
