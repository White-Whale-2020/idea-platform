/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "going" api endpoint
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

func GoingGet(name string) []byte {
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

func GoingPut(r *http.Request) []byte {
	var tmpMbr GoingMemberS
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
