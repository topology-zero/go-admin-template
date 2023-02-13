package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	onceRequestId sync.Once
	requestIdHook *RequestIdHook
)

type RequestIdHook struct {
	requestId string
}

func NewRequestIdHook() *RequestIdHook {
	onceRequestId.Do(func() {
		requestIdHook = &RequestIdHook{}
	})
	return requestIdHook
}

func SetRequestId(requestId string) {
	NewRequestIdHook().requestId = requestId
}

func (h *RequestIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *RequestIdHook) Fire(entry *logrus.Entry) (err error) {
	entry.Data["requestId"] = h.requestId
	return
}
