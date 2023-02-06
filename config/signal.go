package config

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type fnManager struct {
	lock sync.Mutex
	wg   sync.WaitGroup
	fns  []func()
}

var closeManager = new(fnManager)

func init() {
	log.Println("listener close signal")
	go func() {
		var ch = make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

		<-ch

		for _, fn := range closeManager.fns {
			go SafeRun(fn)
		}
	}()

}

// RegisterCloseFn 注册程序退出时需要执行的方法
// 该方法可以调用多次，但是返回的方法只允许调用一次，否则会阻塞线程
func RegisterCloseFn(fn func()) func() {
	closeManager.wg.Add(1)

	closeManager.lock.Lock()
	defer closeManager.lock.Unlock()

	closeManager.fns = append(closeManager.fns, func() {
		defer closeManager.wg.Done()
		fn()
	})
	return closeManager.wg.Wait
}

// SafeRun 安全执行方法
func SafeRun(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%#v\n", r)
		}
	}()

	fn()
}
