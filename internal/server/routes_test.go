package server

import (
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"testing"
)

func TestServer_Routes(t *testing.T) {
	type fields struct {
		app *fiber.App
		db  *pgxpool.Pool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				app: tt.fields.app,
				db:  tt.fields.db,
			}
			s.Routes()
		})
	}
}

func TestServer_RoutesExample(t *testing.T) {
	type fields struct {
		app *fiber.App
		db  *pgxpool.Pool
	}
	type args struct {
		route fiber.Router
		ctrl  *example_domain.ControllerHTTP
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				app: tt.fields.app,
				db:  tt.fields.db,
			}
			s.RoutesExample(tt.args.route, tt.args.ctrl)
		})
	}
}
