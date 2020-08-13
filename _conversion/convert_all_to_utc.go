package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

//"bytes"

//"net/http"

//"time"

type Venue struct {
	Address string `json:"address" bson:"address"`
	Lat     string `json:"lat" bson:"lat"`
	Lng     string `json:"lng" bson:"lng"`
}

type Stats struct {
	Playerid string `json:"playerid" bson:"playerid"`
	Place    string `json:"place" bson:"place"`
	Result   string `json:"result" bson:"result"`
}

// Contest : used to store each event with a list of outcomes, games played and the venue
type Contest struct {
	Start       string   `json:"start" bson:"start"`
	Startoffset string   `json:"startoffset" bson:"startoffset"`
	Stop        string   `json:"stop" bson:"stop"`
	Stopoffset  string   `json:"stopoffset" bson:"stopoffset"`
	Venue       Venue    `json:"venue" bson:"venue"`
	Outcome     []Stats  `json:"outcome" bson:"outcome"`
	Games       []string `json:"games" bson:"games"`
}

type Contests struct {
	Contests []Contest `json:"contests" bson:"contests"`
}

func main() {
	var cc Contests

	// Open our jsonFile
	jsonFile, err := os.Open("stg_records.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened stg_records.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cc)

	for index, _ := range cc.Contests {

		//convert times to UTC
		const longForm = "2006-01-02T15:04:05"
		startfinaltime, err := time.Parse(longForm, cc.Contests[index].Start)
		stopfinaltime, err := time.Parse(longForm, cc.Contests[index].Stop)

		startduration := strings.Replace(cc.Contests[index].Startoffset, ":", "", -1)
		startduration = strings.Replace(startduration, "0", "", -1)

		stopduration := strings.Replace(cc.Contests[index].Stopoffset, ":", "", -1)
		stopduration = strings.Replace(stopduration, "0", "", -1)

		startfinalduration, err := strconv.Atoi(startduration)
		if err != nil {
			fmt.Println(err)
		}
		stopfinalduration, err := strconv.Atoi(stopduration)
		if err != nil {
			fmt.Println(err)
		}

		//startfinaltime = startfinaltime.Add(time.Hour * time.Duration(-startfinalduration)).Format("2006-01-02T15:04:05Z07:00")
		//stopfinaltime = stopfinaltime.Add(time.Hour * time.Duration(-stopfinalduration)).Format("2006-01-02T15:04:05Z07:00")

		cc.Contests[index].Start = startfinaltime.Add(time.Hour * time.Duration(-startfinalduration)).Format("2006-01-02T15:04:05Z07:00")
		cc.Contests[index].Stop = stopfinaltime.Add(time.Hour * time.Duration(-stopfinalduration)).Format("2006-01-02T15:04:05Z07:00")
	}

	fmt.Println("%s ", cc)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(cc)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("stg_records.json", dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
