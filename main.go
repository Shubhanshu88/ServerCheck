package main

import (
	"fmt"
	"net/http"
	"time"
)

var s int

func main() {
	test := []string{
		"http://google.com",
		"https://gawds-leaderboard.herokuapp.com",
		"http://facebook.com",
		"https://github.com",
	}
	c := make(chan string)
	for _, lk := range test {
		s = 0
		go statusCheck(lk, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			statusCheck(link, c)
		}(l)
	}
}
func statusCheck(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("looks like " + link + " might be currently down!")
		s++
		if s < 5 {
			ch <- link
		} else {
			fmt.Println(link + "is currently down!!")
		}
	} else {
		ch <- link
	}
}
