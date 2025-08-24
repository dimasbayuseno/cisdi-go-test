package dbpostgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
)

func TestNewPgx(t *testing.T) {
	tests := []struct {
		name    string
		wantDb  *pgxpool.Pool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDb, err := NewPgx()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPgx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDb, tt.wantDb) {
				t.Errorf("NewPgx() gotDb = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}
