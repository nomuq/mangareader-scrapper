package main

import (
	"github.com/manga-community/mangareader-scrapper/mangareader"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	URL    = "https://www.mangareader.net/naruto/1/"
	SOURCE = "www.mangareader.net"
)

func TestMangaReadGetInfo(t *testing.T) {
	mr := new(mangareader.Mangareader)
	name, issueNumber := mr.GetInfo(URL)

	assert.Equal(t, "naruto", name)
	assert.Equal(t, "1", issueNumber)
}

func TestRetrieveMangaReaderImageLinks(t *testing.T) {
	mr := new(mangareader.Mangareader)

	comic := new(mangareader.Comic)
	comic.URLSource = URL
	comic.Name = "naruto"
	comic.IssueNumber = "1"
	comic.Source = SOURCE

	links, err := mr.RetrieveImageLinks(comic)

	assert.Equal(t, 53, len(links))
	assert.Nil(t, err)
}

func TestSetupMangaReader(t *testing.T) {
	mr := new(mangareader.Mangareader)

	comic := new(mangareader.Comic)
	comic.Name = "naruto"
	comic.IssueNumber = "1"
	comic.URLSource = URL
	comic.Source = SOURCE

	err := mr.Initialize(comic)

	assert.Nil(t, err)
	assert.Equal(t, 53, len(comic.Links))
}

func TestMangareaderRetrieveIssueLinks(t *testing.T) {
	mr := new(mangareader.Mangareader)
	issues, err := mr.RetrieveIssueLinks("https://www.mangareader.net/naruto", false, false)

	assert.Nil(t, err)
	assert.Equal(t, 700, len(issues))
}

func TestMangareaderRetrieveIssueLinksLast(t *testing.T) {
	mr := new(mangareader.Mangareader)
	issues, err := mr.RetrieveIssueLinks("https://www.mangareader.net/naruto", false, true)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(issues))
}

func TestMangaReaderRetrieveLastIssueLink(t *testing.T) {
	mr := new(mangareader.Mangareader)
	issue, err := mr.RetrieveLastIssue("https://www.mangareader.net/naruto")

	assert.Nil(t, err)
	assert.Equal(t, "https://www.mangareader.net/naruto/700", issue)
}
