/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "room" api endpoint
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

type RoomsS struct {
	Endpoint   string  `json:"endpoint"`
	APIversion string  `json:"api-version"`
	Rooms      []RoomS `json:"rooms"`
}

type RoomS struct {
	Nbr     string   `json:"room-nbr"`
	Members []string `json:"members"`
}

func TestRoomJSON(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var rs RoomsS
	json.Unmarshal(rtn, &rs)
	rtn, err = json.Marshal(rs)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}
