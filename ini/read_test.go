package ini_test

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.com/byarbrough/wg-confs/ini"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func TestReadConfig(t *testing.T) {
	buffer := []byte(`
	[Interface]
	PrivateKey = OPRGA+cLdEcPLIPKns/f1rqhzHTxuu1MS4ZNTxxinVc=
	ListenPort = 51820
	FwMark = 0x4D2
	`)

	// Get the configuration as a Config
	got, err := ini.ReadConfig(buffer)
	if err != nil {
		t.Error(err)
	}

	// Construct the expected values
	privateKey, err := wgtypes.ParseKey("OPRGA+cLdEcPLIPKns/f1rqhzHTxuu1MS4ZNTxxinVc=")
	if err != nil {
		log.Fatalf("Unable to read private key: %s", err)
	}

	want := &ini.Config{
		PrivateKey: privateKey,
		ListenPort: uint16(51820),
		FwMark:     uint32(0x4D2),
	}

	if !cmp.Equal(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
