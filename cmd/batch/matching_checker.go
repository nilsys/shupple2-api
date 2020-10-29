package main

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/logger"

	"github.com/urfave/cli"
)

func (b *Batch) cliMatchingChecker() cli.Command {
	return cli.Command{
		Name:   "matching_checker",
		Action: b.matchingChecker,
	}
}

const (
	matchingCheckerTick = 1 * time.Minute
)

func (b *Batch) matchingChecker(c *cli.Context) {
	for range time.Tick(matchingCheckerTick) {
		if err := b.UserCommandRepository.UpdateMatchingExpiredUserLatestMatchingUserID(); err != nil {
			logger.Error(err.Error())
		}
	}
}
