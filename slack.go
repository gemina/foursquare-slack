package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

type Attachment struct {
	Fallback   string  `json:"fallback,omitempty"`
	Color      string  `json:"color,omitempty"`
	AuthorName string  `json:"author_name,omitempty"`
	AuthorLink string  `json:"author_link,omitempty"`
	AuthorIcon string  `json:"author_icon,omitempty"`
	Title      string  `json:"title,omitempty"`
	TitleLink  string  `json:"title_link,omitempty"`
	Pretext    string  `json:"pretext,omitempty"`
	Fields     []Field `json:"fields,omitempty"`
	ImageURL   string  `json:"image_url,omitempty"`
	ThumbURL   string  `json:"thumb_url,omitempty"`
	FooterIcon string  `json:"footer,omitempty"`
	Footer     string  `json:"footer_icon,omitempty"`
	Timestamp  int     `json:"ts,omitempty"`
}

type Message struct {
	Text        string        `json:"text"`
	Channel     string        `json:"channel,omitempty"`
	UserName    string        `json:"username,omitempty"`
	IconURL     string        `json:"icon_url,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

func postToSlack(checkIn *CheckIn) {

	shout := ""
	city := ""
	
	if len(checkIn.Response.Recent[0].Shout) > 0 {
		shout = fmt.Sprintf(" - %s ", checkIn.Response.Recent[0].Shout)
	}
	
	if len(checkIn.Response.Recent[0].Venue.Location.City) > 0 {
		city = fmt.Sprintf("%s, ", checkIn.Response.Recent[0].Venue.Location.City)
	}
	
	text := fmt.Sprintf("%s %s%s\n%s [%s]\n%s%s [<http://maps.google.com/?q=%.6f,%.6f|map>]",
		checkIn.Response.Recent[0].User.FirstName,
		checkIn.Response.Recent[0].User.LastName,
		shout,
		checkIn.Response.Recent[0].Venue.Name,
		checkIn.Response.Recent[0].Venue.Categories[0].Name,
		city,
		checkIn.Response.Recent[0].Venue.Location.Country,
		checkIn.Response.Recent[0].Venue.Location.Lat,
		checkIn.Response.Recent[0].Venue.Location.Lng)

	fallback := fmt.Sprintf("Check-in: %s %s%s @ %s [%s], (%s%s) [<http://maps.google.com/?q=%.6f,%.6f|map>]",
		checkIn.Response.Recent[0].User.FirstName,
		checkIn.Response.Recent[0].User.LastName,
		shout,
		checkIn.Response.Recent[0].Venue.Name,
		checkIn.Response.Recent[0].Venue.Categories[0].Name,
		city,
		checkIn.Response.Recent[0].Venue.Location.Country,
		checkIn.Response.Recent[0].Venue.Location.Lat,
		checkIn.Response.Recent[0].Venue.Location.Lng)

	buf, err := json.Marshal(Message{
		Channel: config.Channel,
		Attachments: []*Attachment{
			{
				Fallback: fallback,
				Color:    "#FFA500",
				Fields: []Field{
					{
						Title: "Check in",
						Value: text,
						Short: false,
					},
				},
			},
		}})

	resp, err := http.Post(config.SlackURL, "application/json", bytes.NewReader(buf))
	checkErr(err)
	defer resp.Body.Close()
}
