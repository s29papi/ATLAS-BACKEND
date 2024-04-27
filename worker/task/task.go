package task

import "time"

type Task interface {
	Duration() time.Duration
	StartProcess(time.Duration)
	StopProcess()
}
