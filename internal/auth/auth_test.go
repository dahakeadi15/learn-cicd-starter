package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	correctHeader := http.Header{}
	correctHeader.Add("Authorization", "ApiKey YWRpdHlhZGFoYWtl")

	incorrectHeaderLen := http.Header{}
	incorrectHeaderLen.Add("Authorization", "ApiKey")

	incorrectAuthHeader := http.Header{}
	incorrectAuthHeader.Add("Authorization", "Bearer dGhpcyBpcyBteSB0b2tlbiEgRG8gbm90IHRvdWNoISE=")

	tests := map[string]struct {
		header  http.Header
		want    string
		wantErr bool
	}{
		"simple": {
			header:  correctHeader,
			want:    "YWRpdHlhZGFoYWtl",
			wantErr: false,
		},
		"noAuthHeader": {
			header:  http.Header{},
			want:    "",
			wantErr: true,
		},
		"incorrectHeaderLength": {
			header:  incorrectHeaderLen,
			want:    "",
			wantErr: true,
		},
		"incorrectAuthHeader": {
			header:  incorrectAuthHeader,
			want:    "",
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("GetAPIKey() got = %v, want %v", got, tc.want)
			}
		})
	}
}
