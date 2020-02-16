package fastproxy

//ProxySettings defines the options for a single proxy
type ProxySettings struct {
	FromHost string `mapstructure:"fromhost"`
	FromPort int    `mapstructure:"fromport"`
	ToHost   string `mapstructure:"tohost"`
	ToPort   int    `mapstructure:"toport"`
}

//Config defines a slice of proxies
type Config struct {
	Proxies []ProxySettings `mapstructure:"proxysettings"`
}
