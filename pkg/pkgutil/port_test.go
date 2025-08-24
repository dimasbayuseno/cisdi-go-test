package pkgutil

import "testing"

func TestGetPort(t *testing.T) {
	type args struct {
		ports []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPort(tt.args.ports...); got != tt.want {
				t.Errorf("GetPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
