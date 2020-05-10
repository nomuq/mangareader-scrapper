package main

import (
	"fmt"
	"github.com/bake/mangadex"
	"github.com/gosimple/slug"
	"github.com/manga-community/mangareader-scrapper/mangareader"
	"github.com/nokusukun/jikan2go/anime"
	"log"
	"strings"
)

func main() {

	md := mangadex.New()
	m, _, err := md.Manga("20068")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Title)

	mr := mangareader.Mangareader{}
	chapters, _ := mr.Latest()
	for _, chapter := range chapters {
		if mr.IsSingleIssue(chapter) {
			title, chapterNumber := mr.GetInfo(chapter)

			isMapped := false

			result, err := anime.Search(anime.Query{Q: title}) // same goes for manga.Search

			if err != nil {
				fmt.Println(err)
			}

			for _, r := range result.Results {
				//fmt.Println(r.MalID, r.Title, title)
				if strings.ToUpper(slug.Make(r.Title)) == strings.ToUpper(title) {
					fmt.Println(r.MalID, r.Title, chapterNumber)
					isMapped = true
					break
				}
			}

			if !isMapped {
				fmt.Println("NOT MAPPED", title)
			}
		}
	}
}
