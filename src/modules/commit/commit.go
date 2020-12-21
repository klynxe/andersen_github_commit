package commit

type Commit struct {
	Comm string `bson:"commit"`
}

func New(commit string) *Commit {
	return &Commit{
		Comm: commit,
	}
}
