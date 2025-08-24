package example_domain

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
	"testing"
)

func TestControllerHTTP_Create(t *testing.T) {
	type fields struct {
		svc Service
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ControllerHTTP{
				svc: tt.fields.svc,
			}
			if err := c.Create(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestControllerHTTP_Delete(t *testing.T) {
	type fields struct {
		svc Service
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ControllerHTTP{
				svc: tt.fields.svc,
			}
			if err := c.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestControllerHTTP_GetByID(t *testing.T) {
	type fields struct {
		svc Service
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ControllerHTTP{
				svc: tt.fields.svc,
			}
			if err := c.GetByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestControllerHTTP_Update(t *testing.T) {
	type fields struct {
		svc Service
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ControllerHTTP{
				svc: tt.fields.svc,
			}
			if err := c.Update(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		svc Service
	}
	tests := []struct {
		name string
		args args
		want *ControllerHTTP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
