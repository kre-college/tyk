package gateway

import (
	"context"
	"github.com/TykTechnologies/tyk/config"
	"github.com/TykTechnologies/tyk/user"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	CreatePolicy(ctx context.Context, policy *user.Policy) error
}

type AnalyticsDB struct {
	*mongo.Database
}

func (db *AnalyticsDB) CreatePolicy(ctx context.Context, user *user.Policy) error {
	_, err := db.Collection(config.TykPolicies).InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(path string) (*AnalyticsDB, error) {
	uri, err := config.CreateMongoDBURI(path)
	if err != nil {
		return nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &AnalyticsDB{
		client.Database(config.TykAnalyticsDB),
	}, nil
}
