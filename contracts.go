package uniongo

type Union interface {
	Execute(WorkerFunc, WorkerData)
}

type WorkerFunc func(WorkerData)

type WorkerData struct {
	InputData  interface{}
	OutputData interface{}
}
