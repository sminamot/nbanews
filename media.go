package nbanews

type Media interface {
	TargetURL() string
	Fetch() error
	ArticleList() []*Article
}

type Article struct {
	Title string
	URL   string
}
