package repository

import (
	"log/slog"
	"webhandler/job/domain"
)

type jobRepository struct {
	jobs []domain.Job
}

func NewInMemoryJobRepository() domain.JobRepository {
	return &jobRepository{}
}

func (r *jobRepository) Schedule(job domain.Job) error {
	r.jobs = append(r.jobs, job)
	slog.Info("job scheduled", "job", job.Code)
	return nil
}
