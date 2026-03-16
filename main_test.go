package main

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct{ name, want string }{
		{"", "Hello, World!"},
		{"Alice", "Hello, Alice!"},
		{"Bob", "Hello, Bob!"},
	}
	for _, tt := range tests {
		got := greeting(tt.name)
		if got != tt.want {
			t.Errorf("greeting(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
