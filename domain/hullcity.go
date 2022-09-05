package domain

import (
	"go.uber.org/zap"
	"time"
)

const DateTime = "2006-01-02 15:04:05"

type NewsletterNewsItem []struct {
	ArticleURL        string `xml:"ArticleURL"`
	NewsArticleID     string `xml:"NewsArticleID"`
	PublishDate       string `xml:"PublishDate"`
	Taxonomies        string `xml:"Taxonomies"`
	TeaserText        string `xml:"TeaserText"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL"`
	Title             string `xml:"Title"`
	OptaMatchId       string `xml:"OptaMatchId"`
	LastUpdateDate    string `xml:"LastUpdateDate"`
	IsPublished       bool   `xml:"IsPublished"`
	BodyText          string `xml:"BodyText"`
	Subtitle          string `xml:"Subtitle"`
}

type ArticleInfo struct {
	ArticleURL        string `xml:"ArticleURL"`
	NewsArticleID     string `xml:"NewsArticleID"`
	PublishDate       string `xml:"PublishDate"`
	Taxonomies        string `xml:"Taxonomies"`
	TeaserText        string `xml:"TeaserText"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL"`
	Title             string `xml:"Title"`
	OptaMatchId       string `xml:"OptaMatchId"`
	LastUpdateDate    string `xml:"LastUpdateDate"`
	IsPublished       bool   `xml:"IsPublished"`
	BodyText          string `xml:"BodyText"`
	Subtitle          string `xml:"Subtitle"`
}

type NewListInformation struct {
	ClubName            string `xml:"ClubName"`
	ClubWebsiteURL      string `xml:"ClubWebsiteURL"`
	NewsletterNewsItems struct {
		NewsletterNewsItem []ArticleInfo `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}

// Adapt Will adapt ArticleInfo to News.
func (n *ArticleInfo) Adapt(logger *zap.Logger, clubName string) News {
	publishedDate, err := time.Parse(DateTime, n.PublishDate)
	if err != nil {
		logger.Error("Error parsing time",
			zap.String("PublishDate", n.PublishDate),
			zap.Error(err),
		)
	}
	news := News{
		ClubName:    clubName,
		Url:         n.ArticleURL,
		Id:          n.NewsArticleID,
		Published:   publishedDate,
		Teaser:      n.TeaserText,
		ImageUrl:    n.ThumbnailImageURL,
		Title:       n.Title,
		OptaMatchId: n.OptaMatchId,
		Content:     n.BodyText,
	}

	return news
}

type NewsArticleInformation struct {
	NewsArticle struct {
		BodyText string `xml:"BodyText"`
		Subtitle string `xml:"Subtitle"`
	} `xml:"NewsArticle"`
}
