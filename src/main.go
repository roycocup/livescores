package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"github.com/julienroland/keyboard-termbox"
	"github.com/buger/goterm"
	term "github.com/nsf/termbox-go"
	"time"
	"github.com/pwaller/keyboard"
	"os"
)

var(
	running = false
	kb keyboard.Keyboard
	c *colly.Collector
)

const (
	SOURCE = "https://www.fifa.com/worldcup/"
)


func main(){
	startTerm()
	bindKeyboard()
	setScraping()
	loop()
}

func startTerm(){
	err := term.Init()
	if err != nil {
		panic(err)
		term.Close()
	}
	running = true

}

func bindKeyboard(){
	kb = termbox.New()
	kb.Bind(func() {
		term.Close()
		running = false
		os.Exit(200)
	}, "space")
}

func setScraping(){
	c = colly.NewCollector(colly.AllowURLRevisit())

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


		fmt.Println()
		fmt.Println(homeName + " vs " + awayName + " - " + matchTime)
		fmt.Println(homeScore + " -- " + awayScore)

	})

}

func listenForExit(){
	event := term.PollEvent()
	if event.Key == term.KeyEsc{
		os.Exit(100)
	}
}

func loop(){

	for running {
		term.Clear(term.ColorGreen, term.ColorBlack)
		term.SetCursor(0,0)
		//goterm.MoveCursor(1,1)
		//kb.Poll(term.PollEvent())
		//listenForExit()

		c.Visit(SOURCE)

		//goterm.Flush()
		term.Flush()

		time.Sleep(time.Second)
	}
}



func extra() {
	c := colly.NewCollector()
	running := true
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	kb := termbox.New()
	kb.Bind(func() {
		fmt.Println("pressed space!")
		running = false
	}, "space")

	goterm.Clear()

	update := 0

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

	for running {
		goterm.MoveCursor(1,1)
		kb.Poll(term.PollEvent())

		c.Visit(SOURCE)
		fmt.Println(update)
		update = update + 1

		goterm.Flush()

		time.Sleep(time.Second)


	}


}
