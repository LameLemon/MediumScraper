package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/labstack/gommon/color"
)

var arguments = struct {
	Input string
}{}

// Article struct hold data scraped from an article
type Article struct {
	// Basic informations
	Title   string
	Summary string

	// Author
	Author Author
}

// Author struct hold data about an author
type Author struct {
	Name string
}

func scrapeArticle(articleLink string) (Article, error) {
	// Create an article structure
	var article Article
	// Create collector
	c := colly.NewCollector()

	// Randomize user agent on every request
	extensions.RandomUserAgent(c)

	// Scrape article's title and summary
	c.OnHTML("div.elevateCover", func(e *colly.HTMLElement) {
		// Scrape username
		article.Title = e.ChildText("h1.elevate-h1")

		// Scrape summary
		article.Summary = e.ChildText("p.elevate-summary")
	})

	// Scrape author informations
	c.OnHTML("div.u-flexEnd", func(e *colly.HTMLElement) {
		// Scrape author name
		article.Author.Name = e.ChildText("a.postMetaInline--author")
	})

	// Visit page and fill collector
	c.Visit(articleLink)

	return article, nil
}

func main() {
	// Parse arguments and fill the arguments structure
	parseArgs(os.Args)

	// Scrape the article
	article, err := scrapeArticle(arguments.Input)
	if err != nil {
		fmt.Println(color.Red("Error while scraping the article: ") + err.Error())
	}

	fmt.Println("Scraping: " + arguments.Input + "\n")
	fmt.Println("Title:   " + article.Title)
	fmt.Println("Summary: " + article.Summary + "\n")
	fmt.Println("Author name: " + article.Author.Name)

}