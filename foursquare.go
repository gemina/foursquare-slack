package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CheckIn struct {
	Meta struct {
		Code        int    `json:"code"`
		RequestID   string `json:"requestId"`
		ErrorDetail string `json:errorDetail`
	} `json:"meta"`
	Response struct {
		Recent []struct {
			ID   string `json:"id"`
			User struct {
				ID        string `json:"id"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
			} `json:"user"`
			Venue struct {
				Name     string `json:"name"`
				Location struct {
					Lat     float64 `json:"lat"`
					Lng     float64 `json:"lng"`
					City    string  `json:"city"`
					Country string  `json:"country"`
				} `json:"location"`
				Categories []struct {
					ID         string `json:"id"`
					Name       string `json:"name"`
					PluralName string `json:"pluralName"`
					ShortName  string `json:"shortName"`
					Icon       struct {
						Prefix string `json:"prefix"`
						Suffix string `json:"suffix"`
					} `json:"icon"`
					Primary bool `json:"primary"`
				} `json:"categories"`
			} `json:"venue"`
		} `json:"recent"`
	} `json:"response"`
}

var lastID string

func getCheckins() {

	var checkIn *CheckIn
	var url = "https://api.foursquare.com/v2/checkins/recent?limit=1&v=20170712"
	url = fmt.Sprintf(url+"&oauth_token=%s", config.OAuth)

	resp, err := http.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &checkIn); err != nil {
		panic(err)
	}

	if checkIn.Meta.Code > 200 {
		fmt.Println(checkIn.Meta.ErrorDetail)
		return
	}

	if len(checkIn.Response.Recent) == 0 {
		return
	}

	if checkIn.Response.Recent[0].ID == lastID {
		return
	}

	lastID = checkIn.Response.Recent[0].ID

	if len(config.IDs) > 0 && !stringInSlice(checkIn.Response.Recent[0].User.ID, config.IDs) {
		return
	}

	go postToSlack(checkIn)
}
