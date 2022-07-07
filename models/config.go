package models

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}
