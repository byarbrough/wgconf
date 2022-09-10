package wgconf_test

import (
	"os"
	"testing"

	"github.com/byarbrough/wgconf"
)

func TestToINI(t *testing.T) {

	interf := new(wgconf.Interface)
	interf.PrivateKey = "2KC6f9xbKYQR1Wsw/X8sRIReHoJJ0B4mBgRBd7Ob3G4="
	interf.ListenPort = 51820
	interf.PreDown = append(interf.PreDown, "first", "second")
	interf.Name = "My first test"

	newPeer := wgconf.Peer{
		PublicKey: "DfNSXkX5tupa3P6VDypiKOhSsb660cHVyr4aNXd2px8=",
	}

	// var peers = []wgconf.Peer{}
	// peers = append(peers, newPeer)

	c := wgconf.Conf{
		Interface: interf,
		Peers:     newPeer,
	}

	got := c.ToINI()

	got.WriteTo(os.Stdout)

}
