package testhelper

import "sync"

type FinishFlagger struct {
	mx         sync.Mutex
	TaskNumber int
	taskDone   int
}

func (f *FinishFlagger) IsFinished() bool {
	f.mx.Lock()
	defer f.mx.Unlock()
	return f.taskDone >= f.TaskNumber
}
func NewFinishFlagger(taskNumber int) *FinishFlagger {
	return &FinishFlagger{TaskNumber: taskNumber}
}
func (f *FinishFlagger) Done() {
	f.mx.Lock()
	f.taskDone++
	f.mx.Unlock()
}
