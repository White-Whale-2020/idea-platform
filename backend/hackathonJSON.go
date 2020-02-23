/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "hackathon" api endpoint
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
	"io/ioutil"
	"log"
)

type HackathonsS struct {
	Endpoint   string            `json:"endpoint"`
	APIversion string            `json:"api-version"`
	Hackathon  []HackathonStruct `json:"hackathons"`
}

type HackathonS struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func TestHackathonJSON(name string) []byte {
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
