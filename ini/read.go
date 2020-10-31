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
//
// See https://git.zx2c4.com/wireguard-tools/about/src/man/wg.8
func ReadConfig(source interface{}) (*Config, error) {

	// Read the INI from source
	cfg, err := ini.ShadowLoad(source)
	if err != nil {
		log.Fatal(err)
	}

	// The new configuration
	config := new(Config)

	// Pull out the conf [Interface] section
	interfaceSection, err := cfg.GetSection("Interface")
	if err != nil {
		return config, fmt.Errorf("Section [Interface] not found: %s", err)
	}

	// PrivateKey is required
	privateKey, err := wgtypes.ParseKey(interfaceSection.Key("PrivateKey").String())
	if err != nil {
		return config, fmt.Errorf("Unable to parse PrivateKey: %s", err)
	}
	config.PrivateKey = privateKey

	// Other keys are optional
	if interfaceSection.HasKey("ListenPort") {
		listenPort, err := strconv.ParseUint(interfaceSection.Key("ListenPort").Value(), 10, 16)
		if err != nil {
			return config, fmt.Errorf("Unable to parse ListenPort: %s", err)
		}
		// Have to go through a function to return a pointer to primitive because nil matters
		config.ListenPort = uint16(listenPort)
	}
	if interfaceSection.HasKey("FwMark") {
		fwMark, err := strconv.ParseUint(interfaceSection.Key("FwMark").Value(), 0, 32)
		if err != nil {
			return config, fmt.Errorf("Unable to parse FwMark: %s", err)
		}
		// Have to go through a function to return a pointer to primitive because nil matters
		config.FwMark = uint32(fwMark)
	}

	return config, err
}
