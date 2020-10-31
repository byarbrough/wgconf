package ini_test

import (
	"log"
	"testing"

	"gitlab.com/byarbrough/wg-confs/ini"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func TestReadWgConf(t *testing.T) {
	buffer := []byte(`
	[Interface]
	PrivateKey = OPRGA+cLdEcPLIPKns/f1rqhzHTxuu1MS4ZNTxxinVc=
	`)

	got, err := ini.ReadWgConf(buffer)
	if err != nil {
		t.Error(err)
	}

	privateKey, err := wgtypes.ParseKey("OPRGA+cLdEcPLIPKns/f1rqhzHTxuu1MS4ZNTxxinVc=")
	if err != nil {
		log.Fatalf("Unable to read private key: %s", err)
	}

	want := wgtypes.Config{
		PrivateKey: &privateKey,
	}

	if got.PrivateKey.String() != want.PrivateKey.String() {
		t.Fatalf("got %v want %v", got, want)
	}
}
