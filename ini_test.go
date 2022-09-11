package wgconf_test

import (
	"os"
	"testing"

	"github.com/byarbrough/wgconf"
)

func TestToINI(t *testing.T) {

	interf := wgconf.Interface{
		PrivateKey: "2KC6f9xbKYQR1Wsw/X8sRIReHoJJ0B4mBgRBd7Ob3G4=",
		ListenPort: 51820,
		PreDown:    []string{"first", "second"},
		Name:       "My first test",
	}

	newPeer := wgconf.Peer{
		PublicKey: "DfNSXkX5tupa3P6VDypiKOhSsb660cHVyr4aNXd2px8=",
		Name:      "First Peer",
	}
	newPeer1 := wgconf.Peer{
		PublicKey: "1ZWjaFuTNEuR4qYua4xN6hLJPhK75CmEiUrXwpoLD1Y=",
	}

	c := wgconf.Conf{
		Interface: interf,
		Peers:     []wgconf.Peer{newPeer, newPeer1},
	}

	c.WriteINI(os.Stdout)

	// got, err := c.ToINI()
	// if err != nil {
	// 	t.Error(err)
	// }

	// got.WriteTo(os.Stdout)

}
