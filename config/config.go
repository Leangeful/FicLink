package config

import(
	"fmt"

	"github.com/tkanos/gonfig"
) 

type config struct{
	CX     	string 
	APIKey 	string
	Version string
}

//Cfg api data
var Cfg = config{}

//ReadConfig gets content of config.json
func ReadConfig(){
	err := gonfig.GetConf("config.json", &Cfg)
	if err != nil{
		fmt.Println("Failed to read config: ", err)
	}
}
