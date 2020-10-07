package main

import (
	"github.com/urfave/cli"
)

func (b *Batch) cliImportSnsShareCount() cli.Command {
	return cli.Command{
		Name: "import_sns_share_count",
		Subcommands: []cli.Command{
			{
				Name:        "twitter",
				Description: "twitterのシェア数集計",
				Action:      b.importTwitterShareCount,
			},
			{
				// facebookはgraph apiにリクエスト制限があるので分けて実行する
				Name:        "facebook",
				Description: "facebookのシェア数集計",
				Subcommands: []cli.Command{
					{
						Name:        "post",
						Description: "Postのシェア数集計",
						Action:      b.importPostFacebookShareCount,
					},
					{
						Name:        "vlog",
						Description: "Vlogのシェア数集計",
						Action:      b.importVlogFacebookShareCount,
					},
					{
						Name:        "cf_project",
						Description: "CfProjectのシェア数集計",
						Action:      b.importCfProjectFacebookShareCount,
					},
				},
			},
		},
	}
}

func (b *Batch) importPostFacebookShareCount(c *cli.Context) error {
	return b.ImportSnsShareCountFacade.ImportPostFacebookShareCount()
}

func (b *Batch) importVlogFacebookShareCount(c *cli.Context) error {
	return b.ImportSnsShareCountFacade.ImportVlogFacebookShareCount()
}

func (b *Batch) importCfProjectFacebookShareCount(c *cli.Context) error {
	return b.ImportSnsShareCountFacade.ImportCfProjectFacebookShareCount()
}

func (b *Batch) importTwitterShareCount(c *cli.Context) error {
	return b.ImportSnsShareCountFacade.ImportTwitterShareCount()
}
