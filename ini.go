package wgconf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func (c *Conf) ToINI() (*ini.File, error) {
	opts := ini.LoadOptions{
		AllowShadows:               true,
		AllowDuplicateShadowValues: true,
	}
	cfg := ini.Empty(opts)

	cfg.NewSection("Interface")
	cfg.Section("Interface").ReflectFrom(&c.Interface)

	for i, peer := range c.Peers {
		thisPeer := fmt.Sprintf("peer_%d", i)
		cfg.NewSection(thisPeer)
		cfg.Section(thisPeer).ReflectFrom(&peer)
	}

	// cfg.WriteTo(os.Stdout)
	// body := cfg.Section("Interface").Body()
	// fmt.Print(body)

	// err := cfg.Append("Interface", interfaceFile)
	// if err != nil {
	// 	return nil, err
	// }

	return cfg, nil

	// Add Required sections
	// cfg.NewSection("Interface")
	// cfg.Section("Interface").Comment = c.Name
	// fmt.Println(c)
	// ini.ReflectFrom(cfg, &c)
	//.Section("Interface").ReflectFrom(c)

	// cfg.Section("Interface").NewKey("PrivateKey", c.PrivateKey)
	// cfg.Section("Interface").NewKey("PresharedKey", "")
	// for _, k := range cfg.Section("Interface").Keys() {
	// 	if k.Value() == "" {
	// 		cfg.Section("Interface").DeleteKey(k.Name())
	// 	}
	// }

	// return cfg
}

func (p *Peer) ToINI() *ini.File {
	return toINI(p)
}

func (interf *Interface) ToINI() *ini.File {
	return toINI(interf)
}

func toINI(myStruct interface{}) *ini.File {
	opts := ini.LoadOptions{
		AllowShadows:               true,
		AllowDuplicateShadowValues: true,
	}
	cfg := ini.Empty(opts)
	ini.ReflectFrom(cfg, myStruct)

	return cfg
}
