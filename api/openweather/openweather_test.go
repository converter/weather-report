package openweather

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComposeRequestURI(t *testing.T) {
	for _, tt := range []struct {
		name      string
		by        string
		arg       string
		wantError bool
	}{
		{"good searchby city", "city", "Fooville,us-pa", false},
		{"good searchby zipcode", "zipcode", "02134", false},
		{"good searchby latlon", "latlon", "32,42", false},
		{"bad searchby citypo", "citypo", "", true},
		{"bad searchby empty string", "", "", true},
		{"bad latlon", "latlon", "32, ", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			uri, err := ComposeRequestURI("DUMMYKEY", tt.by, tt.arg)
			if tt.wantError {
				require.Error(t, err)
				require.Empty(t, uri)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, uri)
			}
		})
	}
}
