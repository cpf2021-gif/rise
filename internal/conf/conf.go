package conf

import (
	"github.com/cpf2021-gif/rise/common/database"
	"github.com/cpf2021-gif/rise/common/net/chttp"
	"go.uber.org/dig"
)

type Conf struct {
	dig.Out

	Server *chttp.Config    `yaml:"server"`
	DB     *database.Config `yaml:"db"`
}
