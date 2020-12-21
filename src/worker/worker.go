package worker

import (
	"andersen/src/modules/commit"
	"andersen/src/modules/setting"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const ParamRequestSince = "since"

var (
	ErrorReadCommits      = errors.New("Error read commits")
	ErrorUnmarshalCommits = "Error unmurshal commits"
)

type Worker struct {
	conf    *Config
	commit  commit.Repository
	setting setting.Repository
}

func NewWorker(commit commit.Repository, setting setting.Repository, conf *Config) (*Worker, error) {
	worker := &Worker{
		conf:    conf,
		commit:  commit,
		setting: setting,
	}
	return worker, nil
}

func (w *Worker) Run() error {
	fmt.Printf("%v Run get commits \n", time.Now().Format("2006-01-02 15:04:05"))
	connect, err := w.setting.GetConnect()
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodGet, w.conf.GetURL(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "token "+w.conf.GetToken())

	q := req.URL.Query()
	q.Add(ParamRequestSince, connect.GetTimeLastConnect().Format(time.RFC3339))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ErrorReadCommits
	}
	commits, err := w.parseCommits(body)
	if err != nil {
		return err
	}

	if len(commits) < 1 {
		fmt.Println("Commits not found")
		return nil
	}

	if err = w.commit.Save(commits); err != nil {
		return err
	}

	connect.SetTimeLastConnect(time.Now())
	if err = w.setting.SaveConnect(connect); err != nil {
		return err
	}

	fmt.Printf("saved %d commits \n", len(commits))
	return nil
}

func (w *Worker) parseCommits(v []byte) ([]*commit.Commit, error) {
	var strCommits []json.RawMessage
	if err := json.Unmarshal(v, &strCommits); err != nil {
		return nil, fmt.Errorf("%v: %w \n %v", ErrorUnmarshalCommits, err, string(v))
	}
	commits := make([]*commit.Commit, 0, len(strCommits))

	for i := range strCommits {
		commits = append(commits, commit.New(string(strCommits[i])))
	}
	return commits, nil
}
