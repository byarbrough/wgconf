package wgconf

// Conf defines the VPN settings for the local node. Cannot use WireGuard's
// "Interface" because that is a reserved keyword in Go.
type Conf struct {
	// Comment in INI syntax used to help keep track of which config
	// section belongs to which node, it's completely ignored by WireGuard
	// and has no effect on VPN behavior.
	Name string `ini:",comment"`
	// Defines what address range the local node should route traffic for.
	Address string `ini:",omitempty"`
	// When the node is acting as a public bounce server, it should hardcode
	// a port to listen for incoming VPN connections from the public internet.
	// Clients not acting as relays should not set this value.
	ListenPort uint16 `ini:",omitempty"`
	// This is the private key for the local node, never shared with other servers.
	// All nodes must have a private key set.
	PrivateKey string `ini:",omitempty"`
	// A shared secret key between all peers. If this is configured, then all peers
	// must have it. Should be randomly generated 32 byte number, base64 encoded.
	// May be generated with same function as PrivateKey
	PresharedKey string `ini:",omitempty"`
	// The DNS server(s) to announce to VPN clients via DHCP,
	// most clients will use this server for DNS requests over the VPN,
	// but clients can also override this value locally on their nodes
	DNS string `ini:",omitempty"`
	// Optionally run a command before the interface is brought up.
	// This option can be specified multiple times, with commands executed in order.
	PreUp []string `ini:",omitempty,allowshadow"`
	// Optionally run a command after the interface is brought up.
	// This option can appear multiple times, as with PreUp.
	PostUp []string `ini:",omitempty,allowshadow"`
	// Optionally run a command before the interface is brought down.
	// This option can appear multiple times, as with PreUp.
	PreDown []string `ini:",omitempty,allowshadow"`
	// Optionally run a command after the interface is brought down.
	// This option can appear multiple times, as with PreUp.
	PostDown []string `ini:",omitempty,allowshadow"`
}

// Defines the VPN settings for a remote peer capable of routing traffic
// for one or more addresses (itself and/or other peers).
// Peers can be either a public bounce server that relays traffic to other peers,
// or a directly accessible client via LAN/internet that is not behind a NAT and
// only routes traffic for itself.
type Peer struct {
	// Comment in INI syntax used to help keep track of which config
	// section belongs to which node, it's completely ignored by WireGuard
	// and has no effect on VPN behavior.
	Name string
	// Defines the publicly accessible address for a remote peer.
	EndPoint string
	// The Ip ranges for which a peer will route traffic
	AllowedIPs string
	// This is the public key for the remote node, shareable with all peers.
	PublicKey string
	// A shared secret key between all peers. If this is configured, then all peers
	// must have it. Should be randomly generated 32 byte number, base64 encoded.
	// May be generated with same function as PrivateKey
	PresharedKey string
	// How many seconds between outgoing pings to send to the peer.
	// Keeps bidirectional connectiosn alive in the NAT router's connection table.
	PersistentKeepalive int
}

// NewConf returns a Conf with a pre-populated private key.
// Also returns the corresponding public key.
func NewConf() (*Conf, string, error) {
	newConf := new(Conf)
	privateKey, err := GenKey()
	newConf.PrivateKey = privateKey
	if err != nil {
		return nil, "", err
	}

	pubKey, err := PubKey(newConf.PrivateKey)
	if err != nil {
		return nil, "", nil
	}

	return newConf, pubKey, nil
}
