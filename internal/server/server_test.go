package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Run(t *testing.T) {
	type fields struct {
		app *fiber.App
		db  *pgxpool.Pool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				app: tt.fields.app,
				db:  tt.fields.db,
			}
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
