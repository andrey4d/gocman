package main

import (
	"testing"
)

func Test_getIdByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIdByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
