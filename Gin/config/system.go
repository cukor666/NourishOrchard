package config

// SystemConfig 系统配置
type SystemConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
}
