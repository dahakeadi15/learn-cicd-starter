package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		key     string
		value   string
		want    string
		wantErr string
	}{
		"noAuthHeader": {
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded.Error(),
		},
		"emptyAuthHeader": {
			key:     "Authorization",
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded.Error(),
		},
		"malformedAuthHeader": {
			key:     "Authorization",
			value:   "-",
			want:    "",
			wantErr: "malformed authorization header",
		},
		"incorrectAuthHeader": {
			key:     "Authorization",
			value:   "Bearer dGhpcyBpcyBteSB0b2tlbiEgRG8gbm90IHRvdWNoISE=",
			want:    "",
			wantErr: "malformed authorization header",
		},
		"incompleteAuthHeader": {
			key:     "Authorization",
			value:   "ApiKey",
			want:    "",
			wantErr: "malformed authorization header",
		},
		"correctAuthHeader": {
			key:     "Authorization",
			value:   "ApiKey YWRpdHlhZGFoYWtl",
			want:    "YWRpdHlhZGFoYWtl",
			wantErr: "not expecting an error",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			got, err := GetAPIKey(header)
			if err != nil {
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				}
			}
			if got != tc.want {
				t.Fatalf("GetAPIKey() got = %v, want %v", got, tc.want)
			}
		})
	}
}
