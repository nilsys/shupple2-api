package main

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/urfave/cli"
)

func (b *Batch) cliImportWordpress() cli.Command {
	return cli.Command{
		Name:  "import_wordpress",
		Usage: "指定したwordpressのリソースをimport/updateする",
		Subcommands: []cli.Command{
			{
				Name:   wordpress.EntityTypePost.String(),
				Usage:  wordpress.EntityTypePost.String(),
				Action: b.importWordpress(wordpress.EntityTypePost),
			},
			{
				Name:   wordpress.EntityTypeLocation.String(),
				Usage:  wordpress.EntityTypeLocation.String(),
				Action: b.importWordpress(wordpress.EntityTypeLocation),
			},
		},
	}
}

func (b *Batch) importWordpress(typ wordpress.EntityType) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		for _, idStr := range c.Args() {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return errors.Wrap(err, "invalid id")
			}

			if err := b.WordpressCallbackService.Import(typ, id); err != nil {
				return errors.Wrapf(err, "failed to import %s(id=%d)", typ, id)
			}
		}

		return nil
	}
}
