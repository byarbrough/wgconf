package wgconf

import (
	"fmt"
	"io"

	"gopkg.in/ini.v1"
)

// WriteINI
func (c *Conf) WriteINI(w io.Writer) {
	opts := ini.LoadOptions{
		AllowShadows:               true,
		AllowDuplicateShadowValues: true,
	}
	cfg := ini.Empty(opts)

	// Write [Interface] first
	cfg.NewSection("Interface")
	cfg.Section("Interface").ReflectFrom(&c.Interface)

	// Handle "Name" comment
	cfg.Section("Interface").DeleteKey("Name")
	if c.Interface.Name != "" {
		fmt.Fprintln(w, "#", c.Interface.Name)
	}
	cfg.WriteTo(w)
	fmt.Fprint(w, "\n")

	// Have to manually loop over peers because INI pkg stores sections
	// as a map, so cannot have duplicate [Peer] sections.
	for _, peer := range c.Peers {
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
}

// func (interf *Interface) writeINI(w io.Writer) {

// }
