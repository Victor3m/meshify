package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type conf struct {
	MySQLConf mySQLConf `json:"mySQL"`
	GoOptions goOptions `json:"goOptions"`
}

type mySQLConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type goOptions struct {
	ServerPort string `json:"serverPort"`
}

func newConf() *conf {
	return &conf{
		MySQLConf: mySQLConf{
			"root",
			"password",
		},
		GoOptions: goOptions{
			"8080",
		},
	}
}

func LoadConf(filepath string) (*conf, error) {
	contents, err := os.ReadFile(filepath)
	conf := newConf()
	if err != nil {
		return nil, fmt.Errorf("there was an error reading the configuration file: %w", err)
	}

	json.Unmarshal(contents, conf)

	return conf, nil
}

func (c *conf) GetServerPort() string {
	return c.GoOptions.ServerPort
}

func (c *conf) GetMySQLUser() string {
	return c.MySQLConf.User
}

func (c *conf) GetMySQLPass() string {
	return c.MySQLConf.Password
}
