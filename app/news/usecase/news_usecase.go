package usecase

import (
	"context"
	"go.uber.org/zap"
	"sports_news/domain"
)

type newsUsecase struct {
	newsRepo domain.NewsRepository
	logger   *zap.Logger
}

// NewNewsUsecase will create new an newsUsecase object representation of domain.NewsUsecase interface.
func NewNewsUsecase(nr domain.NewsRepository, logger *zap.Logger) domain.NewsUsecase {
	return &newsUsecase{
		newsRepo: nr,
		logger:   logger,
	}
}

// GetByID Get news by article id.
func (n *newsUsecase) GetByID(ctx context.Context, id string) (domain.News, error) {
	news, err := n.newsRepo.GetByID(ctx, id)
	if err != nil {
		n.logger.Error("Error on get by id",
			zap.String("id", id), zap.Error(err))
	}
	return news, err
}

// GetByIDAndTeam Get news by article and team id.
func (n *newsUsecase) GetByIDAndTeam(ctx context.Context, id string, teamName string) (domain.News, error) {
	news, err := n.newsRepo.GetByIDAndTeam(ctx, id, teamName)
	if err != nil {
		n.logger.Error("Error on GetByIDAndTeam",
			zap.String("id", id), zap.Error(err))
	}
	return news, err
}

// Fetch Get all news.
func (n *newsUsecase) Fetch(ctx context.Context) ([]domain.News, error) {
	news, err := n.newsRepo.Fetch(ctx)
	if err != nil {
		n.logger.Error("Error on Fetch",
			zap.Error(err))
	}
	return news, err
}

// FetchByTeam Get news by team id.
func (n *newsUsecase) FetchByTeam(ctx context.Context, teamName string) ([]domain.News, error) {
	news, err := n.newsRepo.FetchByTeam(ctx, teamName)
	if err != nil {
		n.logger.Error("Error on Fetch",
			zap.Error(err))
	}
	return news, err
}
