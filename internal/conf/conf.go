package conf

import (
	"github.com/cpf2021-gif/rise/common/database"
	"github.com/cpf2021-gif/rise/common/net/chttp"
)

type Conf struct {
	Server *chttp.Config    `yaml:"server"`
	DB     *database.Config `yaml:"db"`
}
