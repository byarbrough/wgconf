package wgconf

import (
	"gopkg.in/ini.v1"
)

func (c *Conf) ToINI() *ini.File {
	opts := ini.LoadOptions{
		AllowShadows:               true,
		AllowDuplicateShadowValues: true,
	}
	cfg := ini.Empty(opts)

	// Add Required sections
	// cfg.NewSection("Interface")
	// cfg.Section("Interface").Comment = c.Name
	// fmt.Println(c)
	ini.ReflectFrom(cfg, &c)
	//.Section("Interface").ReflectFrom(c)

	// cfg.Section("Interface").NewKey("PrivateKey", c.PrivateKey)
	// cfg.Section("Interface").NewKey("PresharedKey", "")
	// for _, k := range cfg.Section("Interface").Keys() {
	// 	if k.Value() == "" {
	// 		cfg.Section("Interface").DeleteKey(k.Name())
	// 	}
	// }

	return cfg
}
