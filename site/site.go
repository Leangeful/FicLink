package site

import (
	"regexp"
)

var sites = []string{"/s/[0-9]*"}

//Determine which ff site is used
func Determine(url string) int{
	c := 0
	for i, v := range sites {
		r, _ := regexp.Compile(v)
		if r.MatchString(url) {
			return i
		}
		c++
	}
	return -1
}
