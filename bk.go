package nbanews

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type MediaBK struct {
	Articles []*Article
}

func (bk *MediaBK) TargetURL() string {
	return "https://basketballking.jp/news/category/world/nba"
}

func (bk *MediaBK) Fetch() error {
	doc, err := goquery.NewDocument(bk.TargetURL())
	if err != nil {
		return err
	}
	doc.Find("body > div.container > div > div.contents-main > div > div.archives-list").Find("a").Each(func(index int, s *goquery.Selection) {
		u, _ := s.Attr("href")
		a := &Article{
			Title: strings.TrimSpace(s.Find("div.news-category-list__vertical__title").Text()),
			URL:   u,
		}
		bk.Articles = append(bk.Articles, a)
	})

	return nil
}

func (bk *MediaBK) ArticleList() []*Article {
	return bk.Articles
}

func NewMediaBK() Media {
	return &MediaBK{}
}
