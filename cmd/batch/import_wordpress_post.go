package main

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/urfave/cli"
)

func (b *Batch) cliImportWordpressPost() cli.Command {
	return cli.Command{
		Name:  "import_wordpress_post",
		Usage: "指定したwordpressのpostをimport/updateする",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:     flagNameID,
				Usage:    "取り込むpostのID",
				Required: true,
			},
		},
		Action: b.importWordpressPost,
	}
}

func (b *Batch) importWordpressPost(c cli.Context) error {
	return b.WordpressCallbackService.Import(wordpress.EntityTypePost, c.Int(flagNameID))
}
