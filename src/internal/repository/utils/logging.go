package utils

import "src/pkg/logging"

type mockLogger struct{}

func NewMockLogger() logging.Interface {
	return &mockLogger{}
}

func (m *mockLogger) Debugf(message string, args ...interface{}) {}

func (m *mockLogger) Infof(message string, args ...interface{}) {}

func (m *mockLogger) Warnf(message string, args ...interface{}) {}

func (m *mockLogger) Errorf(message string, args ...interface{}) {}

func (m *mockLogger) Fatalf(message string, args ...interface{}) {}
