package main

import (
	"fmt"
	"github.com/manga-community/mangareader-scrapper/mangareader"
)

func main() {
	mr := mangareader.Mangareader{}
	chapters, _ := mr.Latest()
	for _, chapter := range chapters {
		fmt.Println(chapter)
	}
}
