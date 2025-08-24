package migration

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func TestFresh(t *testing.T) {
	tests := []struct {
		name string
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fresh(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fresh() = %v, want %v", got, tt.want)
			}
		})
	}
}
