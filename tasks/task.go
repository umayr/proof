package tasks

import "time"

type Task interface {
	Register(retries uint)
	Schedule(priority uint, time time.Time, data interface{})
}