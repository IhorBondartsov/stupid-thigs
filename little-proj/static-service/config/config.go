package config

import "flag"

type Config struct{
	Port int `json:"port"`
	Host string `json:"host"`
}


func ReadFlags() Config{
	c := Config{}
	c.Port = *flag.Int("port", 8000, "PORT" )
	c.Host = *flag.String("host", "localhost", "Machine host")
	return  c
}
