package wgconf_test

import (
	"testing"

	"github.com/byarbrough/wgconf"
)

func TestNewConf(t *testing.T) {
	newConf, gotPub, err := wgconf.NewInterface()
	if err != nil {
		t.Error(err)
	}

	// check PrivateKey resolves to public key
	wantPub, err := wgconf.PubKey(newConf.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	if gotPub != wantPub {
		t.Errorf("Bad public key, got %s, want %s", gotPub, wantPub)
	}
}
