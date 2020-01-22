package sync

type WaitGroup struct {
	lock Mutex
	n    int
}

func (wg *WaitGroup) Add(n int) {
	wg.lock.Lock()
	defer wg.lock.Unlock()
	if wg.n+n < 0 {
		panic("sync: negative WaitGroup counter")
	}
	wg.n += n
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	panic("sync: unimplemented: WaitGroup.Wait()")
}
