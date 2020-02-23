/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "going" api endpoint
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

type GoingS struct {
	Endpoint     string         `json:"endpoint"`
	APIversion   string         `json:"api-version"`
	GoingMembers []GoingMemberS `json:"members"`
}

type GoingMemberS struct {
	Name  string  `json:"name"`
	Team  string  `json:"team"`
	Forms []FormS `json:forms`
}

type FormS struct {
	Name string `json:"form-name"`
	In   bool   `json:"turned-in"`
}

func TestGoingJSON(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var gs GoingS
	json.Unmarshal(rtn, &gs)
	rtn, err = json.Marshal(gs)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}
