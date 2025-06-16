//go:build unit

package config

import (
	"reflect"
	"testing"
)

func TestGetDBConfig(t *testing.T) {
	tests := []struct {
		name string
		want DBConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDBConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
