package converter

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/config"
)

type (
	Converters struct {
		*config.Config
	}
)

var ConverterSet = wire.NewSet(
	wire.Struct(new(Converters), "*"),
)

func (c Converters) filesURL() config.URL {
	return c.Config.Shupple.FilesURL
}
