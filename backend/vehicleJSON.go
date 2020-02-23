/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "vehicle" api endpoint
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

type VehiclesS struct {
	Endpoint   string     `json:"endpoint"`
	APIversion string     `json:"api-version"`
	Vehicle    []VehicleS `json:"vehicles"`
}

type VehicleS struct {
	Nbr    string   `json:"vehicle=nbr"`
	Driver string   `json:"driver"`
	Nav    string   `json:"navigator"`
	Pass   []string `json:"passengers"`
}

func TestVehicleJSON(name string) []byte {
	rtn, err := ioutil.ReadFile(name)
	if err != nil {
		log.Print(err)
		log.Println("File error")
		return rtn
	}
	var vs VehiclesS
	json.Unmarshal(rtn, &vs)
	rtn, err = json.Marshal(vs)
	if err != nil {
		log.Print(err)
		log.Println("Error remarshaling json")
		return rtn
	}
	return rtn
}
