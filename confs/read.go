package confs

import (
	"fmt"
	"log"
	"net"

	"gopkg.in/ini.v1"
)

// Config is a WireGuard INI representation
// See https://git.zx2c4.com/wireguard-tools/about/src/man/wg.8
type Config struct {
	PrivateKey string
	ListenPort uint16
	FwMark     uint32
	Peers      []*Peer
}

// Peer is a WireGuard [Peer] section
type Peer struct {
	PublicKey           string
	PresharedKey        string
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

	// Pull out the conf [Interface] section
	interfaceSection, err := cfg.GetSection("Interface")
	if err != nil {
		return nil, fmt.Errorf("Section [Interface] not found: %s", err)
	}
	// Read [Interface] into a Config
	config, err := ReadInterface(interfaceSection)
	if err != nil {
		return nil, err
	}

	// Handle any [Peer] sections
	sections := cfg.Sections()

	for s := range sections {
		if sections[s].Name() == "Peer" {
			peer, err := ReadPeer(sections[s])
			if err != nil {
				return config, fmt.Errorf("Error parsing Peer: %s", err)
			}
			// Add the new peer to the config
			config.Peers = append(config.Peers, peer)
		}
	}
	return config, nil

}

// ReadInterface handles only the [Interface] section of a config file
func ReadInterface(section *ini.Section) (*Config, error) {

	// The config that will be returned
	config := new(Config)

	// Verify that section is [Interface]
	if section.Name() != "Interface" {
		return nil, fmt.Errorf("Section must have name \"%s\"", "Interface")
	}

	err := section.StrictMapTo(config)
	if err != nil {
		return nil, fmt.Errorf("Error maping [Interface]: %s", err)
	}

	return config, nil

}

// ReadPeer handles only the [Peer] section of a config file
func ReadPeer(section *ini.Section) (*Peer, error) {

	// The peer that will be returned
	peer := new(Peer)

	// Verify that secion is [Peer]
	if section.Name() != "Peer" {
		return nil, fmt.Errorf("Section must have name \"%s\"", "Peer")
	}

	err := section.StrictMapTo(peer)
	if err != nil {
		return nil, fmt.Errorf("Error maping [Peer]: %s", err)
	}

	return peer, nil
}

// GetSection returns the ini.Section with sectionName from source
func GetSection(source interface{}, sectionName string) (*ini.Section, error) {
	// Read the INI from source
	cfg, err := ini.ShadowLoad(source)
	if err != nil {
		return nil, err
	}

	// Pull out section
	return cfg.GetSection(sectionName)
}
