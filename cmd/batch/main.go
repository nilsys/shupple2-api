package main

import (
	"log"
	"os"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/logger"
	"github.com/uma-co82/shupple2-api/pkg/config"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
	"github.com/urfave/cli"
)

type Batch struct {
	Config *config.Config
	repository.UserCommandRepository
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
	app.Name = "shupple-batch"
	app.HelpName = app.Name
	app.Version = b.Config.Version

	app.Commands = []cli.Command{
		b.cliMatchingChecker(),
	}

	return app.Run(args)
}
