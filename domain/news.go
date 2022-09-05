package domain

import (
	"context"
	"time"
)

type News struct {
	Id          string    `json:"id" bson:"_id"`
	ClubName    string    `json:"teamId" bson:"teamId"`
	OptaMatchId string    `json:"optaMatchId" bson:"optaMatchId,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Type        []string  `json:"type" bson:"type,omitempty"`
	Teaser      string    `json:"teaser" bson:"teaser,omitempty"`
	Content     string    `json:"content" bson:"content,omitempty"`
	Url         string    `json:"url" bson:"url,omitempty"`
	ImageUrl    string    `json:"imageUrl" bson:"imageUrl,omitempty"`
	GalleryUrls []string  `json:"galleryUrls" bson:"galleryUrls,omitempty"`
	VideoUrl    string    `json:"videoUrl" bson:"videoUrl,omitempty"`
	Published   time.Time `json:"published" bson:"published"`
}

// NewsRepository represent the news's repository contract
type NewsRepository interface {
	GetByID(ctx context.Context, id string) (News, error)
	GetByIDAndTeam(ctx context.Context, id string, teamName string) (News, error)
	Fetch(ctx context.Context) ([]News, error)
	FetchByTeam(ctx context.Context, teamName string) ([]News, error)
	Upsert(ctx context.Context, news News) error
}

type NewsUsecase interface {
	GetByID(ctx context.Context, id string) (News, error)
	GetByIDAndTeam(ctx context.Context, id string, teamName string) (News, error)
	Fetch(ctx context.Context) ([]News, error)
	FetchByTeam(ctx context.Context, teamName string) ([]News, error)
}
