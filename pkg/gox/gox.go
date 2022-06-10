package gox

import (
	"fmt"
	"log"
	"runtime"
)

type GoOption struct {
	panicBack func()
}

type GoOptionFunc func(option *GoOption)

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				err = fmt.Errorf("goroutine: panic recovered: %s\n%s", err, buf)
				log.Println(fmt.Sprintf("goroutine: panic recovered: %s\n%s", err, buf))
			}
		}()
		f()
	}()
}
