package main

import (
	"encoding/xml"
)

type AtomFeed struct {
	AtomTag xml.Name    `xml:"feed"`
	Title   string      `xml:"title"`
	Items   []AtomEntry `xml:"entry"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
}

type AtomEntry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	URL     AtomLink `xml:"link"`
}

func (e AtomEntry) EntryTitle() string {
	return e.Title
}

func (e AtomEntry) EntryURL() string {
	return e.URL.Href
}

func AtomFeedItems(rawFeedData []byte) []FeedEntry {
	feed := AtomFeed{}
	xml.Unmarshal(rawFeedData, &feed)
	ret := []FeedEntry{}
	for _, item := range feed.Items {
		ret = append(ret, item)
	}
	return ret
}