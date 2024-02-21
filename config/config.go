package config

type Config struct {
	DB struct {
		Password string `yaml:"password"`
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		DBName string `yaml:"db_name"`
	} `yaml:"db"`
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
}
