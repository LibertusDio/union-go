package uniongo

import "sync"

type GoroutineUnion struct {
}

func NewGoroutineUnion() Union {
	return &GoroutineUnion{}
}

func (u *GoroutineUnion) Execute(worker WorkerFunc, data WorkerData) {
	go func(data WorkerData) {
		worker(data)
	}(data)
}

type LimitedGoroutineUnion struct {
	MaxWorker     int
	Lock          *sync.Mutex
	PendingWorker []chan bool
	WorkerCount   int
}

func NewLimitedGoroutineUnion(max int) Union {
	return &LimitedGoroutineUnion{
		MaxWorker:     max,
		Lock:          new(sync.Mutex),
		PendingWorker: make([]chan bool, 0),
		WorkerCount:   0,
	}
}
func (u *LimitedGoroutineUnion) Execute(worker WorkerFunc, data WorkerData) {
	s := make(chan bool)
	u.PendingWorker = append(u.PendingWorker, s)

	go func(data WorkerData) {
		<-s
		worker(data)
		u.Lock.Lock()
		u.WorkerCount -= 1
		u.Lock.Unlock()
	}(data)

	u.Lock.Lock()
	empty := false
	for !empty {
		if u.WorkerCount < u.MaxWorker && len(u.PendingWorker) > 0 {
			j := u.PendingWorker[0]
			u.WorkerCount += 1
			u.PendingWorker = u.PendingWorker[1:]
			j <- true
		} else {
			empty = true
		}
	}
	u.Lock.Unlock()

}
