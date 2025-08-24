package pkgutil

import (
	"reflect"
	"testing"
)

func TestHTTPResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		Code    int
		Status  string
		Message string
		Data    any
		Errors  []any
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPResponse{
				Code:    tt.fields.Code,
				Status:  tt.fields.Status,
				Message: tt.fields.Message,
				Data:    tt.fields.Data,
				Errors:  tt.fields.Errors,
			}
			got, err := h.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
