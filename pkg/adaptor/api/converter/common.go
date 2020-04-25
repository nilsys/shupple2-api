package converter

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

type (
	Converters struct {
		*config.Config
	}
)

var ConvertersSet = wire.NewSet(
	wire.Struct(new(Converters), "*"),
)

func (c Converters) filesURL() config.URL {
	return c.Config.Stayway.Media.FilesURL
}
