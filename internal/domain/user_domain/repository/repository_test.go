package repository

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	dbpostgres "github.com/dimasbayuseno/cisdi-go-test/pkg/db/postgres"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		db dbpostgres.Queryer
	}
	tests := []struct {
		name string
		args args
		want *Repository
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

func TestRepository_Create(t *testing.T) {
	type fields struct {
		db dbpostgres.Queryer
	}
	type args struct {
		ctx  context.Context
		data entity.Example
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
			r := Repository{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	type fields struct {
		db dbpostgres.Queryer
	}
	type args struct {
		ctx context.Context
		id  string
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
			r := Repository{
				db: tt.fields.db,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetByID(t *testing.T) {
	type fields struct {
		db dbpostgres.Queryer
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData entity.Example
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				db: tt.fields.db,
			}
			gotData, err := r.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetByID() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {
	type fields struct {
		db dbpostgres.Queryer
	}
	type args struct {
		ctx  context.Context
		data entity.Example
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
			r := Repository{
				db: tt.fields.db,
			}
			if err := r.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
