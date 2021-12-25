package main

import (
	"encoding/xml"
)

type RSSFeed struct {
	RSSVersion xml.Name   `xml:"rss"`
	Channel    RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	XMLName xml.Name  `xml:"channel"`
	Title   string    `xml:"title"`
	Items   []RSSItem `xml:"item"`
}

type RSSItem struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	URL     string   `xml:"link"`
}

func (r RSSItem) EntryTitle() string {
	return r.Title
}

func (r RSSItem) EntryURL() string {
	return r.URL
}

func RSSFeedItems(rawFeedData []byte) []FeedEntry {
	feed := RSSFeed{}
	xml.Unmarshal(rawFeedData, &feed)
	ret := []FeedEntry{}
	for _, item := range feed.Channel.Items {
		ret = append(ret, item)
	}
	return ret
}