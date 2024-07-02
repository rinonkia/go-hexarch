package service

import (
	"testing"
)

func Test_extractToken(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "empty string",
			args:    args{""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "hogehoge",
			args:    args{"hogehoge"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Only bearer",
			args:    args{"Bearer"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "valid bearer token",
			args:    args{"Bearer token"},
			want:    "token",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractToken(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
