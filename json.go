package main

import (
	"fmt"
	"strings"
)

type jsonMain struct {
	Main API `json:"main"`
}

type API struct {
	NowPlaying   string `json:"np" db:"np"`
	Listeners    int    `json:"listeners" db:"listeners"`
	BitRate      int    `json:"bitrate" db:"bitrate"`
	IsAFKStream  bool   `json:"isafkstream" db:"isafkstream"`
	IsStreamDesk bool   `json:"isstreamdesk" db:"isstreamdesk"`

	CurrentTime int64 `json:"current"`
	StartTime   int64 `json:"start_time" db:"start_time"`
	EndTime     int64 `json:"end_time" db:"end_time"`

	LastSet    string `json:"lastset" db:"lastset"`
	TrackID    int    `json:"trackid" db:"trackid"`
	Thread     string `json:"thread" db:"thread"`
	Requesting bool   `json:"requesting" db:"requesting"`

	DJName string `json:"djname" db:"djname"`
	DJ     DJApi  `json:"dj" db:""`

	Queue      []ListEntryAPI `json:"queue"`
	LastPlayed []ListEntryAPI `json:"lp"`
	Tags       Tags           `json:"tags" db:"tags"`
}

type Tags []string

func (t *Tags) Scan(src interface{}) error {
	// handle NULL
	if src == nil {
		*t = []string{}
		return nil
	}

	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("unknown type %t", src)
	}

	*t = strings.Split(s, " ")
	return nil
}

type DJApi struct {
	ID          int    `json:"id" db:"djid"`
	Name        string `json:"djname" db:"djname"`
	Description string `json:"djtext" db:"djtext"`
	Image       string `json:"djimage" db:"djimage"`
	Color       string `json:"djcolor" db:"djcolor"`
	Visible     bool   `json:"visible" db:"visible"`
	Priority    int    `json:"priority" db:"priority"`
	ThemeCSS    string `json:"css" db:"css"`
	ThemeID     int    `json:"theme_id" db:"theme_id"`
	Role        string `json:"role" db:"role"`
}

type ListEntryAPI struct {
	Metadata  string `json:"meta" db:"meta"`
	Time      string `json:"time"`
	Type      int    `json:"type" db:"type"`
	Timestamp int64  `json:"timestamp" db:"timestamp"`
}
