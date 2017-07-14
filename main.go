package main

import (
    "encoding/json"
    "io/ioutil"
    "time"
)

type Config struct {
    OAuth    string   `json:"OAuth"`
    SlackURL string   `json:"SlackURL"`
    IDs      []string `json:"IDs"`
    Channel  string   `json:"channel"`
}

var config Config

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func stringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }
    return false
}

func main() {

    configFile, err := ioutil.ReadFile("config.json")
    checkErr(err)
    err = json.Unmarshal(configFile, &config)
    checkErr(err)

    for range time.Tick(time.Second * 30) {
        go getCheckins()
    }
}
