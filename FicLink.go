package main

import (	
	"fmt"
	
	"github.com/Leangeful/FicLink/config"
	
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit" 		
)

type ficLink struct {
	bot reddit.Bot
}

func (r *ficLink) Post(p *reddit.Post) error {
	res := response(p.SelfText)	
	if res != ""{
		return r.bot.Reply(p.Name, res,)
	}
	return nil
}

func (r *ficLink) Comment(c *reddit.Comment) error {
	res := response(c.Body)	
	if res != ""{
		return r.bot.Reply(c.Name, res,)
	}
	return nil	
}

func main(){	
	config.ReadConfig()
	if bot, err := reddit.NewBotFromAgentFile("ficlink.agent", 0); err != nil{		
		fmt.Println("Failed to create bot handle: ", err)
	}else{		
		cfg := graw.Config{Subreddits: []string{"LgeFF"}, SubredditComments: []string{"LgeFF"}}
		handler := &ficLink{bot: bot}
        if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
                fmt.Println("Failed to start graw run: ", err)
        } else {
                fmt.Println("graw run failed: ", wait())
		}
	}
}