package wgconf_test

import (
	"encoding/base64"
	"testing"

	"github.com/byarbrough/wgconf"
)

func TestGenKey(t *testing.T) {
	got, err := wgconf.GenKey()
	if err != nil {
		t.Error(err)
	}
	decodedKey, err := base64.StdEncoding.DecodeString(got)
	if err != nil {
		t.Error(err)
	}
	keyLength := len(decodedKey)
	if keyLength != 32 {
		t.Errorf("Decoded key length must be 32 bytes, got %d for %s", keyLength, got)
	}
}

func TestPubKey(t *testing.T) {

	var testCases = []struct {
		privateKey string
		want       string
		expectErr  bool
	}{
		{"9zw8/YZQofkzUNWAKVzXO3MA1lgPAsaoX4iwnXl0ECI=", "kvn7tp3wXw/x/Km38ETVg+kNd7UdqFwME3EA5QP29Q4=", false},
		{"8Cb11CvA9Zn/7jCw21J/OJJ/IINEzPyQaIRMNUfKMXs=", "1rFrurDYGvT3bAgH8OlDlkCJWpvH/NuEgzZgYIi410M=", false},
		{"NuEgzZgYIi410M=", "", true}, // key too short
	}

	for _, tc := range testCases {
		got, err := wgconf.PubKey(tc.privateKey)
		if err != nil {
			if tc.expectErr {
				break
			} else {
				t.Error(err)
			}
		}
		if got != tc.want {
			t.Errorf("Got %s, want %s", got, tc.want)
		}
	}
}
