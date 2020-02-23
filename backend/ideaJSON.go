/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "idea" api endpoing
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

func TestIdeaJSON(name string) []byte {
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
