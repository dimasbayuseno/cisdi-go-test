package migration

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func TestUp(t *testing.T) {
	tests := []struct {
		name string
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Up(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Up() = %v, want %v", got, tt.want)
			}
		})
	}
}
