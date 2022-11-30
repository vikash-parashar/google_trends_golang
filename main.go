package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/vikashparashar/google_trends_golang/models"
)

var xmldata models.Rss

const (
	url = "https://trends.google.com/trends/trendingsearches/daily/rss?geo=IN"
)

func main() {
	data := readGoogleTreands() // reads xml response of the request that we send to the google trends

	// xml encoding/decoding
	if err := xml.Unmarshal(data, &xmldata); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\t\t\t\t-----------------------------------")
	fmt.Println("\t\t\t\tBelow Results Are Trending Today ")
	fmt.Println("\t\t\t\t-----------------------------------")
	for i, v := range xmldata.Channel.ItemList {
		fmt.Println("-----------------------------------------------------------------------------------------------------------")
		fmt.Printf("Rank #: %d  -  ", i+1)
		fmt.Printf("Search Term : %v  -  ", v.Title)
		fmt.Printf("People Searched Count : %v\n", v.Traffic)
		fmt.Printf("Url To News : %v\n", v.Link)
		fmt.Println("-----------------------------------------------------------------------------------------------------------")
	}

	// json encoding/decoding

	// json_data, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("-----------------------------------")
	// fmt.Println(" Below Results Are Trending Today ")
	// fmt.Println("-----------------------------------")
	// fmt.Println(string(json_data))

}

// get the xml data froom google trend's RSS
func getGoogleTrends(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res, nil
}

// reads xml response of the request that we send to the google trends
func readGoogleTreands() []byte {
	res, err := getGoogleTrends(url) // get the xml data froom google trend's RSS
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}
