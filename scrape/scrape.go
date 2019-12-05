package scrape

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//FfInfo contains all info from ff.net
type FfInfo struct {
	Title, Author, AuthorURL, Stats, Desc string
}

//FFNet gets all info of ff.net
func FFNet(url string) FfInfo {
	var res FfInfo
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	prof := doc.Find("div[id='profile_top']")
	res.Title = prof.Find("b").Text()
	res.Author = prof.Find("a").First().Text()	
	res.AuthorURL,_ = prof.Find("a").First().Attr("href")
	res.AuthorURL = strings.Join([]string{"https://www.fanfiction.net", res.AuthorURL}, "")
	res.Desc = prof.Find("div[class='xcontrast_txt']").First().Text()
	res.Stats = prof.Find("span[class='xgray xcontrast_txt']").First().Text()

	return res
}