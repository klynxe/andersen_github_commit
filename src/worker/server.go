package worker

import (
	"andersen/src/db"
	"andersen/src/modules/commit"
	"andersen/src/modules/setting"
	"fmt"
	"time"

	"go.uber.org/dig"
)

const (
	ErrorBuildContainer = "Error build container"
	ErrorGetConfig      = "Error get config"
	ErrorCreateWorker   = "Error create worker"
	ErrorRunWorker      = "Error run worker"
	ErrorCloseDb        = "Error close DB"
)

func RunServer() {
	container, err := BuildContainer()
	if err != nil {
		fmt.Printf("%v: %v \n", ErrorBuildContainer, err)
		return
	}
	var config *Config
	err = container.Invoke(func(conf *Config) {
		config = conf

	})
	if err != nil {
		fmt.Printf("%v: %v \n", ErrorGetConfig, err)
		return
	}

	err = container.Invoke(func(worker *Worker) {
		for {
			errRun := worker.Run()
			if errRun != nil {
				fmt.Printf("%v: %v \n", ErrorRunWorker, errRun)
				return
			}
			timer := time.NewTimer(time.Duration(config.GetPeriod()) * time.Second)
			<-timer.C
		}

	})
	if err != nil {
		fmt.Printf("%v: %v \n", ErrorCreateWorker, err)
		return
	}
	defer func() {
		if err := container.Invoke(func(db *db.Db) {
			db.Close()
		}); err != nil {
			fmt.Printf("%v: %v \n", ErrorCloseDb, err)
		}
	}()
}

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(db.NewConfig); err != nil {
		return nil, err
	}
	if err := container.Provide(db.NewDb); err != nil {
		return nil, err
	}
	if err := container.Provide(setting.NewRepository); err != nil {
		return nil, err
	}
	if err := container.Provide(commit.NewRepository); err != nil {
		return nil, err
	}
	if err := container.Provide(NewConfig); err != nil {
		return nil, err
	}
	if err := container.Provide(NewWorker); err != nil {
		return nil, err
	}

	return container, nil
}
