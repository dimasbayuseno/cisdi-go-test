package migration

import (
	"github.com/dimasbayuseno/cisdi-go-test/migration"
	"reflect"
	"testing"
)

func Test_initMigration(t *testing.T) {
	tests := []struct {
		name    string
		want    *migration.Migration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initMigration()
			if (err != nil) != tt.wantErr {
				t.Errorf("initMigration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initMigration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
