package entity

import (
	"testing"
	"time"
)

func TestExample_TableName(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Description string
		Type        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := Example{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Type:        tt.fields.Type,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if got := ex.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
