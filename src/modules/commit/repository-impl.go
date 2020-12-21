package commit

import (
	"andersen/src/db"
	"context"
)

const MONGO_COLLECTION = "github"

var _ Repository = (*RepositoryImpl)(nil)

type RepositoryImpl struct {
	db *db.Db
}

func NewRepository(db *db.Db) Repository {
	impl := &RepositoryImpl{
		db: db,
	}
	return impl
}

func (impl *RepositoryImpl) Save(commits []*Commit) error {
	if len(commits) < 1 {
		return nil
	}

	collection := impl.db.GetCollection(MONGO_COLLECTION)
	sliceCommits := make([]interface{}, 0, len(commits))
	for i := range commits {
		sliceCommits = append(sliceCommits, *commits[i])

	}
	if _, err := collection.InsertMany(context.Background(), sliceCommits); err != nil {
		return err
	}
	return nil
}
