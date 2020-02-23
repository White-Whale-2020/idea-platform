/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "idea" api endpoing
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

type IdeasS struct {
	Endpoint   string  `json:"endpoint"`
	APIversion string  `json:"api-version"`
	Ideas      []IdeaS `json:"ideas"`
}

type IdeaS struct {
	Name     string `json:"name"`
	Desc     string `json:"description"`
	Env      string `json:"environment"`
	Proposer string `json:"proposer"`
	Done     bool   `json:"done"`
	Date     string `json:"date-proposed"`
}

func IdeaGet(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var is IdeaS
	json.Unmarshal(rtn, &is)
	rtn, err = json.Marshal(is)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}

func IdeaPut(r *http.Request) []byte {
	var tmpMbr IdeaS
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
