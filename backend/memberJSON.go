/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "member" api endpoint
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

type UsersS struct {
	Endpoint   string  `json:"endpoint"`
	APIversion string  `json:"api-version"`
	Users      []UserS `json:"users"`
}

type UserS struct {
	Name   string   `json:"name"`
	Major  string   `json:"major"`
	Skills []string `json:"skills"`
}

func TestMemberJSON(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var us UsersS
	json.Unmarshal(rtn, &us)
	rtn, err = json.Marshal(us)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}
