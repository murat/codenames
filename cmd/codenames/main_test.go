package main

import (
	"path/filepath"
	"runtime"
	"testing"
)

func Test_run(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "happy path",
			args:    []string{filepath.Join(filepath.Dir(b), "../../words/en.txt")},
			wantErr: false,
		},
		{
			name:    "unhappy path",
			args:    []string{"not_exist.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
