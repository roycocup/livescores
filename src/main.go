package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

const (
	SOURCE = "https://www.fifa.com/worldcup/"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})

	//type item struct {
	//	Team1 string,
	//	Team2 string,
	//}

	selector := "div.fi-mu__item > a.fi-mu__link > div > div.fi-mu__m"

	c.OnHTML(selector, func(e *colly.HTMLElement) {

		homeTeam := "div.home"
		homeName := e.ChildText(homeTeam + " > div > span.fi-t__nText")

		if homeName == "{HOMETEAMNAME}" {
			return
		}

		awayTeam := "div.away"
		awayName := e.ChildText(awayTeam + " > div > span.fi-t__nText")

		info := "div.fi-mu__score-info"
		matchTime := e.ChildText(info + "> div.fi-mu__match-time > span")

		score := info + " > div.fi-mu__score-wrap"
		homeScore := e.ChildText(score + "> div.home")
		awayScore := e.ChildText(score + "> div.away")

		fmt.Println(homeName + " vs " + awayName + " - " + matchTime)
		fmt.Println(homeScore + " -- " + awayScore)

		fmt.Println()

	})

	c.Visit(SOURCE)
}
