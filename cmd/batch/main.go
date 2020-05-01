package main

import (
	"log"
	"os"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/urfave/cli"
)

const (
	flagNameID    = "id"
	flagNameMedia = "media"
	flagNameSpan  = "span"
)

type Batch struct {
	Config                   *config.Config
	WordpressCallbackService service.WordpressCallbackService
	PostQueryRepository      repository.PostQueryRepository
	PostCommandRepository    repository.PostCommandRepository
	ReviewQueryRepository    repository.ReviewQueryRepository
	ReviewCommandRepository  repository.ReviewCommandRepository
	VlogQueryRepository      repository.VlogQueryRepository
	VlogCommandRepository    repository.VlogCommandRepository
	FeatureQueryRepository   repository.FeatureQueryRepository
	FeatureCommandRepository repository.FeatureCommandRepository
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
		b.cliImportViews(),
	}

	return app.Run(args)
}
