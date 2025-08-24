package exception

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestFiberErrorHandler(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FiberErrorHandler(tt.args.ctx, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("FiberErrorHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPanicIfNeeded(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PanicIfNeeded(tt.args.err)
		})
	}
}
