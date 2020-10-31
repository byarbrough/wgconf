package ini

import (
	"log"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"gopkg.in/ini.v1"
)

// ReadWgConf read a wireguard conf file and return a pointer to the Config
func ReadWgConf(source interface{}) (*wgtypes.Config, error) {

	cfg, err := ini.ShadowLoad(source)
	if err != nil {
		log.Fatal(err)
	}

	// Break out key=value
	privateKey, err := wgtypes.ParseKey(cfg.Section("Interface").Key("PrivateKey").String())
	if err != nil {
		log.Fatalf("Unable to read private key: %s", err)
	}

	// Place values in to config
	config := new(wgtypes.Config)
	config.PrivateKey = &privateKey

	return config, nil
}
