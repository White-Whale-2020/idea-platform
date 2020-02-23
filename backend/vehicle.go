/*
	Application: 	IDP (IDea Platform)
	File:					idp.go
	Description:	source file for "vehicle" api endpoint
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

func VehicleGet(name string) []byte {
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

func VehiclePut(r *http.Request) []byte {
	var tmpMbr VehicleS
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
