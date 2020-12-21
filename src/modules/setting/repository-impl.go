package setting

import (
	"andersen/src/db"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrorGetConnect       = errors.New("Error get connect")
	ErrorUnmarshalConnect = "Error unmarshal connect"
)

const MONGO_COLLECTION = "setting"

var _ Repository = (*RepositoryImpl)(nil)

type RepositoryImpl struct {
	collection *mongo.Collection
}

func NewRepository(db *db.Db) Repository {
	impl := &RepositoryImpl{
		collection: db.GetCollection(MONGO_COLLECTION),
	}
	return impl
}

func (impl *RepositoryImpl) SaveConnect(connect *Connect) error {
	update := bson.M{
		"$set": New(SettingConnectName, connect),
	}
	if _, err := impl.collection.UpdateOne(context.Background(),
		NewFilter(SettingConnectName),
		update,
		options.Update().SetUpsert(true),
	); err != nil {
		return err
	}
	return nil
}

func (impl *RepositoryImpl) GetConnect() (*Connect, error) {
	setting := Setting{}
	if err := impl.collection.FindOne(context.Background(), NewFilter(SettingConnectName)).Decode(&setting); err != nil {
		fmt.Println(err)
		return NewConnect(time.Time{}), nil
	}
	var connect Connect
	jsonString, _ := json.Marshal(setting.GetValue().(bson.D).Map())
	if err := json.Unmarshal(jsonString, &connect); err != nil {
		return nil, fmt.Errorf("%v: %w", ErrorUnmarshalConnect, err)
	}
	return &connect, nil
}
