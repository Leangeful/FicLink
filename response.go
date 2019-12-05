package main

import( 
	"fmt"
	"time"
	"regexp"

	"github.com/Leangeful/FicLink/config"
	"github.com/Leangeful/FicLink/scrape"
	"github.com/Leangeful/FicLink/google"
	"github.com/Leangeful/FicLink/site"
)

var validCmd = regexp.MustCompile(`\!ficlink\((.*)\)`)
var ffnTemplate = "***[%s](%s)*** by *[%s](%s)* \n \n>%s \n \n^(%s)\n\n---\n\n"

func response(body string)string{
	res := ""
	cmds := validCmd.FindAllStringSubmatch(body, 30)
	falseCmdCount := 0
	if cmds != nil {
		<-time.After(10 * time.Second)
		for _, q := range cmds {				
			url := google.Search("Harry Potter Fanfiction " + q[1])			
			site := site.Determine(url)
			switch site {
			case -1:
				falseCmdCount++
			case 0:
				story := scrape.FFNet(url)
				res += fmt.Sprintf(ffnTemplate, story.Title, url, story.Author, 
					story.AuthorURL, story.Desc, story.Stats)				
			}			
		}
		if falseCmdCount != len(cmds){
			fmt.Println("Return Response")
			return res + "**FicLink**^(" + config.Cfg.Version + ")"
		}
	}
return ""
}

