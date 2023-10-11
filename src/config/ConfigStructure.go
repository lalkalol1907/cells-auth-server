package config

type Config struct {
	DB struct {
		Url string `yaml:"url"`
	} `yaml:"db"`
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"redis"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}
