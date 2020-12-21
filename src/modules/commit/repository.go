package commit

type Repository interface {
	Save(commit []*Commit) error
}
