package hullcity

import (
	"context"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"sports_news/app/utils"
	"sports_news/config"
	"sports_news/domain"
	"strconv"
)

type hullCityConsumer struct {
	httpClient http.Client
	logger     *zap.Logger
	newsRepo   domain.NewsRepository
	urls       config.HullCityConfig
}

// NewHullCityConsumer will create new an hullCityConsumer object representation of domain.NewsConsumer interface.
func NewHullCityConsumer(hc http.Client, l *zap.Logger, newsRepo domain.NewsRepository, urls config.HullCityConfig) domain.NewsConsumer {
	return &hullCityConsumer{
		httpClient: hc,
		logger:     l,
		newsRepo:   newsRepo,
		urls:       urls,
	}
}

// Fetch Request news from hull city website.
func (h *hullCityConsumer) Fetch(ctx context.Context, quantity int) {
	data := h.RequestData(quantity)
	newListInformationData, err := utils.ReadXml(data, domain.NewListInformation{})
	if err != nil {
		h.logger.Error("error on ReadXml",
			zap.String("domain", "NewListInformation"),
			zap.Error(err))
		return
	}
	for _, value := range newListInformationData.NewsletterNewsItems.NewsletterNewsItem {
		articleBytes := h.RequestArticle(value.NewsArticleID)
		article, err := utils.ReadXml(articleBytes, domain.NewsArticleInformation{})
		if err != nil {
			h.logger.Error("error on ReadXml",
				zap.String("domain", "NewsArticleInformation"),
				zap.Error(err))
		}
		value.BodyText = article.NewsArticle.BodyText
		value.Subtitle = article.NewsArticle.Subtitle
		err = h.newsRepo.Upsert(ctx, value.Adapt(h.logger, newListInformationData.ClubName))
		if err != nil {
			h.logger.Error("error on upsert",
				zap.String("news id", value.NewsArticleID),
				zap.Error(err))
		}
	}
}

// RequestData Handle the http request to get data from hull city website.
func (h *hullCityConsumer) RequestData(quantity int) []byte {
	url := h.urls.BaseUrl + h.urls.FetchPath + strconv.Itoa(quantity)
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		h.logger.Info("failed to RequestData "+err.Error(),
			zap.String("url", url),
			zap.String("consumer", "hullcity"),
			zap.Int("Step", 1),
			zap.Error(err),
		)
		return nil
	}
	res, err := h.httpClient.Do(req)
	if err != nil {
		h.logger.Info("failed to RequestData "+err.Error(),
			zap.String("url", url),
			zap.String("consumer", "hullcity"),
			zap.Int("Step", 2),
			zap.Error(err),
		)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		h.logger.Info("failed to RequestData "+err.Error(),
			zap.String("url", url),
			zap.String("consumer", "hullcity"),
			zap.Int("Step", 3),
			zap.Error(err),
		)
		return nil
	}
	return body
}

// RequestArticle Handle request to get details from an article from hull city website.
func (h *hullCityConsumer) RequestArticle(id string) []byte {
	url := h.urls.BaseUrl + h.urls.GetIdPath + id
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		h.logger.Info("failed to RequestData "+err.Error(),
			zap.String("url", url),
			zap.String("consumer", "hullcity"),
			zap.Int("Step", 1),
			zap.Error(err),
		)
		return nil
	}
	res, err := h.httpClient.Do(req)
	if err != nil {
		h.logger.Info("failed to RequestData "+err.Error(),
			zap.String("url", url),
			zap.String("consumer", "hullcity"),
			zap.Int("Step", 2),
			zap.Error(err),
		)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		h.logger.Info("failed to RequestData "+err.Error(),
			zap.String("url", url),
			zap.String("consumer", "hullcity"),
			zap.Int("Step", 3),
			zap.Error(err),
		)
		return nil
	}
	return body
}
