package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port string `envconfig:"STOCK_SERVER_PORT"`
	}

	Mysql struct {
		Host     []string `envconfig:"STOCK_MYSQL_HOST"`
		Port     string   `envconfig:"STOCK_MYSQL_PORT"`
		Database string   `envconfig:"STOCK_MYSQL_DATABASE"`
		User     string   `envconfig:"STOCK_MYSQL_USER"`
		Password string   `envconfig:"STOCK_MYSQL_PASSWORD"`
	}
}

var cfg *Config

// configType := os.Args[1] //dev, test, pro // default = dev
func Start(ct string) *Config {
	if cfg == nil {
		c := Config{}
		if ct == "dev" { //dev
			if err := c.readDev(); err != nil {
				os.Exit(1)
			}
		} else if ct == "test" { //test
			if err := c.readTest(); err != nil {
				os.Exit(1)
			}
		} else if ct == "pro" { //pro
			if err := c.readPro(); err != nil {
				os.Exit(1)
			}
		}
		cfg = &c
	}
	return cfg
}

func (c *Config) readDev() error {
	err := envconfig.Process("STOCK", c)
	if err != nil {
		log.Fatal("Failed to load configuration, error is - " + err.Error())
		return err
	}

	//Server Config
	if c.Server.Port == "" {
		c.Server.Port = "3030"
	}

	//MySql Config
	if c.Mysql.Host == nil {
		c.Mysql.Host = []string{"93.187.203.194"}
	}
	if c.Mysql.Port == "" {
		c.Mysql.Port = "3306"
	}
	if c.Mysql.Database == "" {
		c.Mysql.Database = "mynet"
	}
	if c.Mysql.User == "" {
		c.Mysql.User = "mynet"
	}
	if c.Mysql.Password == "" {
		c.Mysql.Password = "#871mtkV"
	}
	return nil
}

func (c *Config) readTest() error {
	err := envconfig.Process("STOCK", c)
	if err != nil {
		log.Fatal("Failed to load configuration, error is - " + err.Error())
		return err
	}

	//Server Config
	if c.Server.Port == "" {
		c.Server.Port = "3030"
	}

	//MySql Config
	if c.Mysql.Host == nil {
		c.Mysql.Host = []string{"93.187.203.194"}
	}
	if c.Mysql.Port == "" {
		c.Mysql.Port = "3306"
	}
	if c.Mysql.Database == "" {
		c.Mysql.Database = "mynet"
	}
	if c.Mysql.User == "" {
		c.Mysql.User = "mynet"
	}
	if c.Mysql.Password == "" {
		c.Mysql.Password = "#871mtkV"
	}
	return nil
}

func (c *Config) readPro() error {
	err := envconfig.Process("STOCK", c)
	if err != nil {
		log.Fatal("Failed to load configuration, error is - " + err.Error())
		return err
	}

	//Server Config
	if c.Server.Port == "" {
		c.Server.Port = "3030"
	}

	//MySql Config
	if c.Mysql.Host == nil {
		c.Mysql.Host = []string{"93.187.203.194"}
	}
	if c.Mysql.Port == "" {
		c.Mysql.Port = "3306"
	}
	if c.Mysql.Database == "" {
		c.Mysql.Database = "mynet"
	}
	if c.Mysql.User == "" {
		c.Mysql.User = "mynet"
	}
	if c.Mysql.Password == "" {
		c.Mysql.Password = "#871mtkV"
	}
	return nil
}
