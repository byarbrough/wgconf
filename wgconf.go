package wgconf

import "net"

// Config is a WireGuard INI representation
// See https://git.zx2c4.com/wireguard-tools/about/src/man/wg.8
type Config struct {
	PrivateKey string
	ListenPort uint16  `ini:",omitempty"`
	FwMark     uint32  `ini:",omitempty"`
	Peers      []*Peer `ini:",omitempty"`
}

// Peer is a WireGuard [Peer] section
type Peer struct {
	PublicKey           string
	PresharedKey        string      `ini:"value,omitempty"`
	AllowedIPs          []net.IPNet `ini:"value,omitempty,allowshadow"`
	EndPoint            string      `ini:"value,omitempty"`
	PersistentKeepalive uint16      `ini:"value,omitempty"`
}
