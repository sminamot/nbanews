package nbanews

import (
	"github.com/PuerkitoBio/goquery"
)

type MediaRT struct {
	Articles []*Article
}

func (rt *MediaRT) TargetURL() string {
	return "https://nba.rakuten.co.jp/news/article/"
}

func (rt *MediaRT) Fetch() error {
	doc, err := goquery.NewDocument(rt.TargetURL())
	if err != nil {
		return err
	}
	doc.Find("body > div.wrapper > main > div > div").Find("article").Each(func(index int, s *goquery.Selection) {
		u, _ := s.Find("a").Attr("href")
		a := &Article{
			Title: s.Find("h1").Text(),
			URL:   u,
		}
		rt.Articles = append(rt.Articles, a)
	})

	return nil
}

func (rt *MediaRT) ArticleList() []*Article {
	return rt.Articles
}

func NewMediaRT() Media {
	return &MediaRT{}
}
