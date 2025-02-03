package domain

import "time"

type Job struct {
	ID           int
	Method       string
	Payload      string
	Code         string
	URL          string
	StartMessage string
	EndMessage   string
	Attempts     int
	CreatedAt    time.Time
	CompletedAt  *time.Time
	Done         bool
}

type JobRepository interface {
	Schedule(job Job) error
}

func NewJob(
	method string,
	payload string,
	code string,
	url string,
	startMessage string,
	endMessage string,
) Job {
	// validate
	return Job{
		ID:           0,
		Method:       method,
		Payload:      payload,
		Code:         code,
		URL:          url,
		StartMessage: startMessage,
		EndMessage:   endMessage,
		CreatedAt:    time.Now(),
		Attempts:     0,
		CompletedAt:  nil,
		Done:         false,
	}
}
