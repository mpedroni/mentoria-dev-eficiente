package handler

type ScheduleJobRequest struct {
	Method       string `json:"method" validate:"required,oneof=GET POST"`
	Payload      string `json:"payload"`
	Code         string `json:"code" validate:"required"`
	URL          string `json:"url" validate:"required,url"`
	StartMessage string `json:"start_message" validate:"required"`
	EndMessage   string `json:"end_message" validate:"required"`
}
