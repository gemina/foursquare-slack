package main

import (
    "bytes"
    "encoding/json"
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

func postToSlack(text string) {

    buf, err := json.Marshal(Message{
        Channel: config.Channel,
        Attachments: []*Attachment{
            {
                Fallback: text,
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
