package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/gocolly/colly"
)

type Anime struct {
	Rank string
	Title string
	Score string
	Type string
	Members string 
}

var Titles = []string{}
var Rank = []string{}
var Score = []string{}
var Type = []string{}
var Members = []string{}

func check(e error) {
	/* 
	Genera un panic en caso de haber un error.
	Based on https://gobyexample.com/writing-files
	*/
    if e != nil {
        panic(e)
    }
}

func main() {

    file, err := os.Create("data.csv")
    check(err)

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Rank", "Title", "Score", "Type", "Members"}
	writer.Write(headers)

	c := colly.NewCollector(
			colly.AllowedDomains("myanimelist.net"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(".rank.ac", func(e *colly.HTMLElement) { // Rank
		rank := e.ChildText(".top-anime-rank-text")
		Rank = append(Rank, rank)
	})

	c.OnHTML(".anime_ranking_h3", func(e *colly.HTMLElement) { // Title
		title := e.ChildText("a")
		Titles = append(Titles, title)
	})

	c.OnHTML(".js-top-ranking-score-col", func(e *colly.HTMLElement) { // Score
		score := e.ChildText(".score-label")
		Score = append(Score, score)
	})

	c.OnHTML(".detail", func(e *colly.HTMLElement) { // Type and Members
		info := e.ChildText(".information")
		split := strings.Split(info, "\n")
		Type = append(Type, strings.Split(split[0], " ")[0])
		Members = append(Members, strings.Replace(strings.TrimSpace(strings.Split(split[2], "members")[0]), ",", "", -1))
	})

	c.OnHTML(".link-blue-box.next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		limit, err := strconv.Atoi(strings.Split(nextPage, "=")[1])
		check(err)
		if limit <= 150 { // Tomaremos los 200 animes con mejor ranking
			c.Visit(nextPage)
		}
	})

	startUrl := "https://myanimelist.net/topanime.php"
	c.Visit(startUrl)

	// Write data in CSV
	for index, value := range Rank {
        anime := Anime{}
		anime.Rank = value
		anime.Title =  Titles[index]
		anime.Score = Score[index]
		anime.Type = Type[index]
		anime.Members = Members[index]
		row := []string{anime.Rank, anime.Title, anime.Score, anime.Type, anime.Members}
		writer.Write(row)
    }


}