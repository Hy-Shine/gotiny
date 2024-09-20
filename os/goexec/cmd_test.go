// go:build linux || darwin
package goexec

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		command string
		opts    []Option
		wantErr bool
	}{
		{"valid command", "ls", nil, false},
		{"invalid command", "non-existent-command", nil, true},
		{"with options", "ls", []Option{WithArg("-l")}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.command, tt.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunOutput(t *testing.T) {
	tests := []struct {
		name    string
		command string
		opts    []Option
		wantErr bool
	}{
		{"valid command", "ls", nil, false},
		{"valid command with options", "ls", []Option{WithArg("-l")}, false},
		{"invalid command", "non-existent-command", nil, true},
		{"command with error output", "ls /non-existent-dir", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := RunOutput(tt.command, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			assert.NotNil(t, output)
		})
	}
}

func TestRunOutput_Output(t *testing.T) {
	output, err := RunOutput("ls", WithArg("-l"))
	assert.Nil(t, err)
	assert.True(t, bytes.Contains(output, []byte("total")))
}
