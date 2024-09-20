//go:build linux || darwin
// +build linux darwin

package goexec

import "testing"

func TestLookPath(t *testing.T) {
	tests := []struct {
		name   string
		exists bool
	}{
		{name: "ls", exists: true},
		{name: "mkdir", exists: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lp := NewLookPath(tt.name)
			lp.LookPath()
			if got := lp.Exists(); got != tt.exists {
				t.Errorf("LookPath() = %v, want %v", got, tt.exists)
			}
		})
	}
}
