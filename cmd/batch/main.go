package main

import (
	"log"
	"os"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/urfave/cli"
)

const (
	flagNameID = "id"
)

type Batch struct {
	Config                   *config.Config
	WordpressCallbackService service.WordpressCallbackService
}

func main() {
	batch, err := InitializeBatch(config.DefaultConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := batch.Run(os.Args); err != nil {
		logger.Fatal(err.Error())
	}
}

func (b *Batch) Run(args []string) error {
	app := cli.NewApp()
	app.Name = "stayway-media-batch"
	app.HelpName = app.Name
	app.Version = b.Config.Version

	app.Commands = []cli.Command{
		b.cliImportWordpressPost(),
	}

	return app.Run(args)
}
