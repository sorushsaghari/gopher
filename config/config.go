package config


import "fmt"

type Config struct {
	Port uint
	Host string
	DbName string
	Dialect string
	Password string
	User string
}

func (c* Config) GetConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", c.Host, c.Port, c.User, c.DbName, c.Password)
}