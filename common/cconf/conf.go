package cconf

import (
	"os"

	"github.com/cpf2021-gif/rise/internal/conf"
	"gopkg.in/yaml.v3"
)

func Unmarshal(filePath string, out interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}

func NewConfig() conf.Conf {
	c := conf.Conf{}
	if err := Unmarshal("etc/config.yaml", &c); err != nil {
		panic(err)
	}

	return c
}
