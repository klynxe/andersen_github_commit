package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ErrCreateDatabase = "Error create database"
	ErrPingDatabase   = "Error ping database"
)

type Db struct {
	conf   *Config
	client *mongo.Client
}

func NewDb(conf *Config) (*Db, error) {
	db := &Db{
		conf: conf,
	}
	if err := db.Open(); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *Db) Open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(db.conf.Url))
	if err != nil {
		return fmt.Errorf("%v: %w", ErrCreateDatabase, err)
	}
	db.client = client

	if err = client.Connect(context.Background()); err != nil {
		return fmt.Errorf("%v: %w", ErrCreateDatabase, err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return fmt.Errorf("%v: %w", ErrPingDatabase, err)
	}
	return nil
}

func (db *Db) Close() error {
	return db.client.Disconnect(context.Background())
}

func (db *Db) GetCollection(collection string) *mongo.Collection {
	return db.client.Database(db.conf.DbName).Collection(collection)
}
