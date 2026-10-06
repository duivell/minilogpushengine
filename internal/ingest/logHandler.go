package ingest

import (
	"mini-log-push/internal/queue"
	"net/http"

	"golang.org/x/time/rate"
)

type logsHandler struct {
	queue   *queue.Queue
	limiter map[string]*rate.Limiter
}

func NewLogHandler() *logsHandler {
	q := queue.New(1000)
	return &logsHandler{queue: q}
}

func (handler *logsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (handler *logsHandler) pushRecords(records []string) {

}
