package main

import (
	"andersen/src/worker"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go worker.RunServer()

	waitStop()
}

func waitStop() {
	stop := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		stop <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated by '%v'", <-stop)
}
