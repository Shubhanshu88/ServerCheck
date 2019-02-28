package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
)

var s int

func main() {

	test := []string{}
	urls := []string{}

	app := cli.NewApp()
	app.Name = "urlcheck"
	app.Version = "1.0.0"
	app.Usage = "A utility to check if specified server is running "
	app.UsageText = "urlcheck [command] [urls ...]"
	app.Author = "Shubhanshu Gairola"
	Flags := []cli.Flag{
		cli.StringFlag{
			Name:  "https",
			Value: "false",
			Usage: "To make https requests",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "check",
			Usage: "To check if specified servers are running",
			Flags: Flags,
			Action: func(c *cli.Context) error {
				links := c.Args()
				if links != nil {
					for _, url := range links {
						urls = append(urls, url)
					}
					test = append(test, urls...)
				} else if links == nil {
					cli.ShowAppHelp(c)
				}
				linkFind(test)
				fmt.Println(test)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func linkFind(test []string) {
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
		fmt.Println(link + " is currently up and running")
		ch <- link
	}
}
