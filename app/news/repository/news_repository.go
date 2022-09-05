package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sports_news/domain"
)

const newsCollectionName = "news"

type mongoNewsRepo struct {
	DB         *mongo.Client
	collection *mongo.Collection
}

// NewMongoNewsRepository will create an implementation of news.Repository.
func NewMongoNewsRepository(db *mongo.Client, dbName string) domain.NewsRepository {
	return &mongoNewsRepo{
		DB:         db,
		collection: db.Database(dbName).Collection(newsCollectionName),
	}
}

// GetByID Get the news by article id.
func (m *mongoNewsRepo) GetByID(ctx context.Context, id string) (domain.News, error) {
	filter := bson.D{{"_id", id}}
	var result domain.News
	err := m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, domain.ErrNotFound
		}
		return result, domain.ErrInternalServerError
	}

	return result, nil
}

// GetByIDAndTeam Get the news by article id and team id.
func (m *mongoNewsRepo) GetByIDAndTeam(ctx context.Context, id string, teamName string) (domain.News, error) {
	filter := bson.D{{"_id", id}, {"teamId", teamName}}
	var result domain.News
	err := m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, domain.ErrNotFound
		}
		return result, domain.ErrInternalServerError
	}

	return result, nil
}

// Fetch Get all news.
func (m *mongoNewsRepo) Fetch(ctx context.Context) ([]domain.News, error) {
	filter := bson.D{{}}
	var results []domain.News
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return results, domain.ErrNotFound
		}
		return results, domain.ErrInternalServerError
	}

	if err = cursor.All(ctx, &results); err != nil {
		return results, domain.ErrInternalServerError
	}

	return results, nil
}

// FetchByTeam Get the news by team id.
func (m *mongoNewsRepo) FetchByTeam(ctx context.Context, teamName string) ([]domain.News, error) {
	filter := bson.D{{"teamId", teamName}}
	var results []domain.News
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return results, domain.ErrNotFound
		}
		return results, domain.ErrInternalServerError
	}

	if err = cursor.All(ctx, &results); err != nil {
		return results, domain.ErrInternalServerError
	}

	if results == nil {
		return results, domain.ErrNotFound
	}

	return results, nil
}

// Upsert Update or insert news, this is a mongodb option on update.
func (m *mongoNewsRepo) Upsert(ctx context.Context, news domain.News) error {
	filter := bson.D{{"_id", news.Id}}
	update := bson.D{{"$set", news}}
	opts := options.Update().SetUpsert(true)
	_, err := m.collection.UpdateOne(ctx, filter, update, opts)

	return err
}
