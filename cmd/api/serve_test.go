package api

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func TestServe(t *testing.T) {
	tests := []struct {
		name string
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Serve(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serve() = %v, want %v", got, tt.want)
			}
		})
	}
}
