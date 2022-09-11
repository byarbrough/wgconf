package wgconf

import (
	"fmt"
	"io"

	"gopkg.in/ini.v1"
)

// WriteINI writes an INI configuration to be used in a WireGuard .conf file
func (c *Conf) WriteINI(w io.Writer) {

	c.Interface.writeINI(w)

	// Have to manually loop over peers because INI pkg stores sections
	// as a map, so cannot have duplicate [Peer] sections.
	for _, peer := range c.Peers {
		peer.writeINI(w)
	}
}

// WriteINI writes an INI configuration to be used in a WireGuard .conf file
func (interf *Interface) WriteINI(w io.Writer) {

	interf.writeINI(w)

}

// WriteINI writes an INI configuration to be used in a WireGuard .conf file
func (peer *Peer) WriteINI(w io.Writer) {

	peer.writeINI(w)

}

func (interf *Interface) writeINI(w io.Writer) {
	opts := ini.LoadOptions{
		AllowShadows:               true,
		AllowDuplicateShadowValues: true,
	}
	cfg := ini.Empty(opts)

	cfg.NewSection("Interface")
	cfg.Section("Interface").ReflectFrom(&interf)

	// Handle "Name" comment
	cfg.Section("Interface").DeleteKey("Name")
	if interf.Name != "" {
		fmt.Fprintln(w, "#", interf.Name)
	}
	cfg.WriteTo(w)
	fmt.Fprint(w, "\n")
}

func (peer *Peer) writeINI(w io.Writer) {
	opts := ini.LoadOptions{
		AllowShadows:               true,
		AllowDuplicateShadowValues: true,
	}
	cfg := ini.Empty(opts)
	cfg.NewSection("Peer")
	cfg.Section("Peer").ReflectFrom(&peer)

	// Handle "Name" comment
	cfg.Section("Peer").DeleteKey("Name")
	if peer.Name != "" {
		fmt.Fprintln(w, "#", peer.Name)
	}
	cfg.WriteTo(w)
	fmt.Fprint(w, "\n")
}
