package ini

import (
	"fmt"
	"log"
	"strconv"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"gopkg.in/ini.v1"
)

// Config is a WireGuard INI representation
type Config struct {
	PrivateKey wgtypes.Key
	ListenPort uint16
	FwMark     uint32
	// Peers      []Peer
}

// Peer is a WireGuard [Peer] section
type Peer struct {
	PublicKey wgtypes.Key
}

// ReadConfig parses WireGuard configuration and returns interface
func ReadConfig(source interface{}) (*Config, error) {

	cfg, err := ini.ShadowLoad(source)
	if err != nil {
		log.Fatal(err)
	}

	// Pull out the INI interface section
	interfaceSection := cfg.Section("Interface")
	// Get values
	privateKey, err := wgtypes.ParseKey(interfaceSection.Key("PrivateKey").String())
	if err != nil {
		return new(Config), fmt.Errorf("Unable to parse PrivateKey: %s", err)
	}
	listenPort, err := strconv.ParseUint(interfaceSection.Key("ListenPort").Value(), 10, 16)
	if err != nil {
		return new(Config), fmt.Errorf("Unable to parse ListenPort: %s", err)
	}
	fwMark, err := strconv.ParseUint(interfaceSection.Key("FwMark").Value(), 0, 32)
	if err != nil {
		return new(Config), fmt.Errorf("Unable to parse FwMark: %s", err)
	}

	Config := Config{
		PrivateKey: privateKey,
		ListenPort: uint16(listenPort),
		FwMark:     uint32(fwMark),
	}

	return &Config, err
}
