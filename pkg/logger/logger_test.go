package logger

import (
	"context"
	"github.com/rs/zerolog"
	"reflect"
	"testing"
)

func TestLog(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *zerolog.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Log(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}
