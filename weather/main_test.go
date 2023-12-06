package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		apiKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.args.apiKey); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_formatAndPrintJson(t *testing.T) {
	type args struct {
		jsonData []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatAndPrintJson(tt.args.jsonData)
		})
	}
}
