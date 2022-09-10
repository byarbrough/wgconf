package wgconf_test

import (
	"os"
	"testing"

	"github.com/byarbrough/wgconf"
)

func TestToINI(t *testing.T) {

	c := new(wgconf.Conf)
	c.PrivateKey = "2KC6f9xbKYQR1Wsw/X8sRIReHoJJ0B4mBgRBd7Ob3G4="
	c.ListenPort = 51820
	c.PreDown = append(c.PreDown, "first", "second")
	c.Name = "My first test"
	got := c.ToINI()

	got.WriteTo(os.Stdout)

}
