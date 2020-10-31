package ini

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"gopkg.in/ini.v1"
)

// Config is a WireGuard INI representation
// See https://git.zx2c4.com/wireguard-tools/about/src/man/wg.8
type Config struct {
	PrivateKey wgtypes.Key
	ListenPort uint16
	FwMark     uint32
	// Peers      []Peer
}

// Peer is a WireGuard [Peer] section
type Peer struct {
	PublicKey           wgtypes.Key
	PresharedKey        wgtypes.Key
	AllowedIPs          []net.IPNet
	EndPoint            string
	PersistentKeepalive uint16
}

// ReadConfig parses WireGuard configuration and returns interface
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
	err = readInterface(config, interfaceSection)

	return config, err
}

// readInterface handles only the [Interface] section of a config file
func readInterface(config *Config, section *ini.Section) error {
	// PrivateKey is required
	privateKey, err := wgtypes.ParseKey(section.Key("PrivateKey").String())
	if err != nil {
		return fmt.Errorf("Unable to parse PrivateKey: %s", err)
	}
	config.PrivateKey = privateKey

	// Other keys are optional
	if section.HasKey("ListenPort") {
		listenPort, err := strconv.ParseUint(section.Key("ListenPort").Value(), 10, 16)
		if err != nil {
			return fmt.Errorf("Unable to parse ListenPort: %s", err)
		}
		// Have to go through a function to return a pointer to primitive because nil matters
		config.ListenPort = uint16(listenPort)
	}
	if section.HasKey("FwMark") {
		fwMark, err := strconv.ParseUint(section.Key("FwMark").Value(), 0, 32)
		if err != nil {
			return fmt.Errorf("Unable to parse FwMark: %s", err)
		}
		// Have to go through a function to return a pointer to primitive because nil matters
		config.FwMark = uint32(fwMark)
	}

	return nil
}
