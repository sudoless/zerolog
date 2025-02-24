//go:build !binary_log && !windows
// +build !binary_log,!windows

package zlog

import (
	"errors"
	"io"
	"testing"
)

var writeCalls int

type mockedWriter struct {
	wantErr bool
}

func (c mockedWriter) Write(p []byte) (int, error) {
	writeCalls++

	if c.wantErr {
		return -1, errors.New("expected error")
	}

	return len(p), nil
}

// Tests that a new writer is only used if it actually works.
func TestResilientMultiWriter(t *testing.T) {
	tests := []struct {
		name    string
		writers []io.Writer
	}{
		{
			name: "All valid writers",
			writers: []io.Writer{
				mockedWriter{
					wantErr: false,
				},
				mockedWriter{
					wantErr: false,
				},
			},
		},
		{
			name: "All invalid writers",
			writers: []io.Writer{
				mockedWriter{
					wantErr: true,
				},
				mockedWriter{
					wantErr: true,
				},
			},
		},
		{
			name: "First invalid writer",
			writers: []io.Writer{
				mockedWriter{
					wantErr: true,
				},
				mockedWriter{
					wantErr: false,
				},
			},
		},
		{
			name: "First valid writer",
			writers: []io.Writer{
				mockedWriter{
					wantErr: false,
				},
				mockedWriter{
					wantErr: true,
				},
			},
		},
	}

	for _, tt := range tests {
		writers := tt.writers
		multiWriter := MultiLevelWriter(writers...)

		logger := New(multiWriter).With().Timestamp().Logger().Level(InfoLevel)
		logger.Info().Msg("Test msg")

		if len(writers) != writeCalls {
			t.Errorf("Expected %d writers to have been called but only %d were.", len(writers), writeCalls)
		}
		writeCalls = 0
	}
}
