package models

import "encoding/xml"

type Rss struct {
	XMlName xml.Name `xml:"rss" json:"rss"`
	Channel *Channel `xml:"channel" json:"channel"`
}

type Channel struct {
	Title    string `xml:"title" json:"title"`
	ItemList []Item `xml:"item" json:"" json:"item"`
}

type Item struct {
	Title      string `xml:"title" json:"title"`
	Link       string `xml:"link" json:"link"`
	Traffic    string `xml:"approx_traffic" json:"approx_traffic"`
	News_Items []News `xml:"news_item" json:"news_item"`
}

type News struct {
	Headline      string `xml:"news_item_title" json:"today_headline"`
	Headline_Link string `xml:"news_item_url" json:"headlines_link"`
}
