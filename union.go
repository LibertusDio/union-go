package uniongo

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
