package handler

import (
	"encoding/json"
	"net/http"
	"webhandler/job/domain"

	"github.com/go-playground/validator/v10"
)

type jobsHandler struct {
	validator *validator.Validate
	repo      domain.JobRepository
}

func New(v *validator.Validate, r domain.JobRepository) *jobsHandler {
	return &jobsHandler{
		validator: v,
		repo:      r,
	}
}

func (h *jobsHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /jobs", h.ScheduleJob)
}

func (h *jobsHandler) ScheduleJob(w http.ResponseWriter, r *http.Request) {
	var body ScheduleJobRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validator.Struct(body); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	job := domain.NewJob(
		body.Method,
		body.Payload,
		body.Code,
		body.URL,
		body.StartMessage,
		body.EndMessage,
	)

	if err := h.repo.Schedule(job); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func errorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
