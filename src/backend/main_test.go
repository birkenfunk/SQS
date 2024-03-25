package main

import "testing"

func TestGetHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"TestGetHelloWorld", "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHelloWorld(); got != tt.want {
				t.Errorf("GetHelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}
