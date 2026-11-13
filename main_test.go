package main

import (
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	for _, tt := range []struct {
		args  []string
		want  time.Duration
		valid bool
	}{
		{[]string{"30s"}, 30 * time.Second, true},
		{[]string{"2m"}, 2 * time.Minute, true},
		{nil, 0, false},
		{[]string{"0s"}, 0, false},
		{[]string{"-1s"}, 0, false},
		{[]string{"5"}, 0, false},
		{[]string{"1s", "2s"}, 0, false},
	} {
		got, err := interval(tt.args)
		if (err == nil) != tt.valid || got != tt.want {
			t.Errorf("interval(%v) = %v, %v; want %v, valid=%v", tt.args, got, err, tt.want, tt.valid)
		}
	}
}
