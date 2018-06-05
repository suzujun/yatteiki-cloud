package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const prefix = "api"

var params Config

type (
	// DbConfig database config struct
	DbConfig struct {
		Host            string
		User            string
		Password        string
		Dbname          string
		Port            string
		MaxIdleConns    int
		ConnMaxLifetime time.Duration
	}
	Config struct {
		// database - master
		DbMasterHost     string `envconfig:"db_master_host" default:"localhost"`
		DbMasterUser     string `envconfig:"db_master_user" default:"yatteiki_admin"`
		DbMasterPassword string `envconfig:"db_master_password" default:"xkMhzFyZWL5ndR0KnACb3KMsKCwbx46n"`
		DbMasterDbname   string `envconfig:"db_master_dbname" default:"yatteiki"`
		DbMasterPort     string `envconfig:"db_master_port" default:"3306"`

		// database slave
		DbSlaveHost     string `envconfig:"db_slave_host" default:"localhost"`
		DbSlaveUser     string `envconfig:"db_slave_user" default:"yatteiki_admin"`
		DbSlavePassword string `envconfig:"db_slave_password" default:"xkMhzFyZWL5ndR0KnACb3KMsKCwbx46n"`
		DbSlaveDbname   string `envconfig:"db_slave_dbname" default:"yatteiki"`
		DbSlavePort     string `envconfig:"db_slave_port" default:"3306"`
	}
)

func Init() error {
	return envconfig.Process(prefix, &params)
}

func Get() Config {
	return params
}

func (c Config) DbMasterConfig() *DbConfig {
	return &DbConfig{
		Host:     c.DbMasterHost,
		User:     c.DbMasterUser,
		Password: c.DbMasterPassword,
		Dbname:   c.DbMasterDbname,
		Port:     c.DbMasterPort,
	}
}

func (c Config) DbSlaveConfig() *DbConfig {
	return &DbConfig{
		Host:     c.DbSlaveHost,
		User:     c.DbSlaveUser,
		Password: c.DbSlavePassword,
		Dbname:   c.DbSlaveDbname,
		Port:     c.DbSlavePort,
	}
}
