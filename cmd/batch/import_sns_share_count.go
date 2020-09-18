package main

import (
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
)

func (b *Batch) cliImportSnsShareCount() cli.Command {
	return cli.Command{
		Name:   "import_sns_share_count",
		Action: b.importSnsShareCount,
	}
}

func (b *Batch) importSnsShareCount(c *cli.Context) error {
	eg := errgroup.Group{}
	eg.Go(b.ImportSnsShareCountFacade.ImportPostSnsShareCount)
	eg.Go(b.ImportSnsShareCountFacade.ImportVlogSnsShareCount)
	eg.Go(b.ImportSnsShareCountFacade.ImportCfProjectSnsShareCount)
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
