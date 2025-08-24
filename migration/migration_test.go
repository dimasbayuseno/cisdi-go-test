package migration

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
)

func TestMigration_Down(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
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
			m := &Migration{
				db: tt.fields.db,
			}
			if err := m.Down(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Down() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigration_Fresh(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
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
			m := &Migration{
				db: tt.fields.db,
			}
			if err := m.Fresh(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Fresh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigration_Up(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
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
			m := &Migration{
				db: tt.fields.db,
			}
			if err := m.Up(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Up() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *Migration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
