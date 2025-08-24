package config

import (
	"github.com/spf13/viper"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name string
		args args
		want *config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.filenames...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetViper(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name string
		args args
		want *viper.Viper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetViper(tt.args.filenames...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetViper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name    string
		args    args
		want    *viper.Viper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.filenames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseConfig(t *testing.T) {
	type args struct {
		v *viper.Viper
	}
	tests := []struct {
		name    string
		args    args
		want    *config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfig(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_database_GetDSN(t *testing.T) {
	type fields struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
		SSLMode  string
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
			d := database{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Name:     tt.fields.Name,
				SSLMode:  tt.fields.SSLMode,
			}
			if got := d.GetDSN(); got != tt.want {
				t.Errorf("GetDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initDefaultValue(t *testing.T) {
	type args struct {
		v *viper.Viper
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initDefaultValue(tt.args.v)
		})
	}
}
