package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			want:    "test-api-key",
			wantErr: false,
		},
		{
			name:    "missing authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key"},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}
