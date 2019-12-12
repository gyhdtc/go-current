package main

import (
	"sync"
	"time"
)

producer := func(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	for i:=5;i>0;i-- {
		l.lock()
		l.Unlock()
		time.Sleep(1)
	}
}

observer := func(wg *sync.WaitGroup, L sync.Locker) {
	defer wg.Done()
	l.lock()
	defer l.Unlock()
}

test := func(count int, mutex, rwmutex sync.Locker)time.Duration {
	var wg sync.WaitGroup
	wg.Add(count+1)
	begintesttime := time.Now()
	go producer(&wg, mutex)
	for i := count; i > 0; i-- {
		go observer(&wg, rwmutex)
	}
	wg.Wait()
	return time.Since(begintesttime)
}

tw := tabwriter.NewWriter(os.stdout, 0, 1, 2, '', 0)
defer tw.Flush()

