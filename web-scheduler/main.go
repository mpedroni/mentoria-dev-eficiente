package main

import (
	"log/slog"
	"net/http"
	"os"
	"webhandler/auth"
	"webhandler/job/handler"
	"webhandler/job/repository"

	"github.com/go-playground/validator/v10"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request received",
			"method", r.Method,
			"url", r.URL.String(),
		)
		next.ServeHTTP(w, r)
	})
}

func MockAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth.SetClientID(ctx, "mock-client-id")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	validator := validator.New()

	jobRepository := repository.NewInMemoryJobRepository()

	handler.
		New(validator, jobRepository).
		Register(mux)

	http.ListenAndServe(":8080", LoggingMiddleware(mux))
}
