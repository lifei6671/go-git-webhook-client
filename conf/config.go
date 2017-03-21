package conf

import "github.com/go-ini/ini"

func GetString(name string,defval string) string {

	cfg, err := ini.LooseLoad("conf/app.conf");
	if err != nil {
		return defval
	}
	section,err := cfg.GetSection("");

	if err != nil {
		return defval
	}
	key,err := section.GetKey(name)

	if err != nil {
		return defval
	}
	return key.String()
}

