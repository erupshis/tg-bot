package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var singleton closer

type callback func()
type closer struct {
	cbs []callback
}

func Add(cb callback) {
	singleton.cbs = append(singleton.cbs, cb)
}

func Run() {
	idleConnsClosed := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh

		wg := sync.WaitGroup{}
		for _, cb := range singleton.cbs {
			wg.Add(1)
			go func(cb callback) {
				defer wg.Done()
				cb()
			}(cb)
		}

		wg.Wait()
		close(idleConnsClosed)
	}()

	<-idleConnsClosed
	log.Printf("service gracefull shutdown done")
}
