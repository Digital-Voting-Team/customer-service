package resources

import "strings"

type AccessLevel int64

const (
	Guest AccessLevel = iota + 1
	Worker
	Manager
	Accountant
	Admin
)

type WorkerStatus string

const (
	Busy      WorkerStatus = "busy"
	Available WorkerStatus = "available"
	Vacation  WorkerStatus = "vacation"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (w *WorkerStatus) Validate(candidate string) bool {
	correct := []string{"busy", "available", "vacation"}
	return contains(correct, strings.ToLower(candidate))
}
