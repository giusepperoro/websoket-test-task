package config

type ServiceConfiguration struct {
	WebsocketRequest string `yaml:"websocketRequest"`
	Ch               string `yaml:"ch"`
	Op               string `yaml:"op"`
}
