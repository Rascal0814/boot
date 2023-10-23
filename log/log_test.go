package log

import (
	"github.com/pkg/errors"
	"testing"
)

func TestError(t *testing.T) {
	logger := NewLogger("boot")

	logger.Error("test", errors.New("test error"))
	logger.Debug("test debug")
	logger.Info("test info", "key", "value")
	logger.Warn("test info", "id", "12")
}
