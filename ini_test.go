package wgconf_test

import (
	"os"
	"testing"

	"github.com/byarbrough/wgconf"
)

func TestWriteINI(t *testing.T) {

	interf := wgconf.Interface{
		PrivateKey: "2KC6f9xbKYQR1Wsw/X8sRIReHoJJ0B4mBgRBd7Ob3G4=",
		ListenPort: 51820,
		PreDown:    []string{"first rule", "second rule"},
		Name:       "My first test",
	}

	newPeer := wgconf.Peer{
		PublicKey: "DfNSXkX5tupa3P6VDypiKOhSsb660cHVyr4aNXd2px8=",
		Name:      "First Peer",
	}
	newPeer1 := wgconf.Peer{
		PublicKey:    "1ZWjaFuTNEuR4qYua4xN6hLJPhK75CmEiUrXwpoLD1Y=",
		PresharedKey: "IFazrJcuMZi9PECth++BfxIY2GdjOWMzqTuQp5Ddb1c=",
	}

	c := wgconf.Conf{
		Interface: interf,
		Peers:     []wgconf.Peer{newPeer, newPeer1},
	}

	t.Run("full conf", func(t *testing.T) {
		c.WriteINI(os.Stderr)
	})

}
