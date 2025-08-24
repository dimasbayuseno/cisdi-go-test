package middleware

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	type args struct {
		duration time.Duration
		opt      []OptFunc
	}
	tests := []struct {
		name string
		args args
		want fiber.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Timeout(tt.args.duration, tt.args.opt...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithExcludePaths(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want OptFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithExcludePaths(tt.args.paths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithExcludePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
