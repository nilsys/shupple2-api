package main

import (
	"log"

	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	payjpRepo "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type Config struct {
	AddSupporterComment struct {
		TargetCfReturnGiftID int    `yaml:"target_cf_return_gift_id"`
		CardNumber           string `yaml:"card_number"`
		CardExpireMonth      string `yaml:"card_expire_month"`
		CardExpireYear       string `yaml:"card_expire_year"`
		Comments             []struct {
			UserID int    `yaml:"user_id"`
			Body   string `yaml:"body"`
		} `yaml:"comments"`
	} `yaml:"add_supporter_comment"`
}

type Script struct {
	// DB                    *gorm.DB
	Config *config.Config
	repository.CfReturnGiftQueryRepository
	repository.UserQueryRepository
	payjpRepo.CustomerQueryRepository
	payjpRepo.CustomerCommandRepository
	service.CardCommandService
	service.ChargeCommandService
	service.ShippingCommandService
	PayJpClient *payjp.Service `wire:"-"`
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "failed to initialize script")
	}

	return script.Run()
}

func (s Script) Run() error {
	if !s.Config.IsDev() {
		return errors.New("this script is only for dev env")
	}

	s.PayJpClient = buildPayjpService(s.Config)
	var configWrapper Config
	if err := s.Config.Scripts.Decode(&configWrapper); err != nil {
		return errors.Wrap(err, "failed to load script config")
	}
	config := configWrapper.AddSupporterComment

	returnGift, err := s.CfReturnGiftQueryRepository.FindByID(config.TargetCfReturnGiftID)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, comment := range config.Comments {
		user, err := s.UserQueryRepository.FindByID(comment.UserID)
		if err != nil {
			return errors.WithStack(err)
		}

		if _, err := s.CustomerQueryRepository.FindCustomer(user.PayjpCustomerID()); err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.WithStack(err)
			}

			if err := s.CustomerCommandRepository.StoreCustomer(user.PayjpCustomerID(), user.Email); err != nil {
				return errors.WithStack(err)
			}
		}

		if err := s.ShippingCommandService.StoreShippingAddress(user, &entity.ShippingAddress{UserID: user.ID}); err != nil {
			return errors.WithStack(err)
		}

		tokenRequest := payjp.Card{
			Number:   config.CardNumber,
			ExpMonth: config.CardExpireMonth,
			ExpYear:  config.CardExpireYear,
		}
		token, err := s.PayJpClient.Token.Create(tokenRequest)
		if err != nil {
		}

		if err := s.CardCommandService.Register(user, token.ID); err != nil {
			if !serror.IsErrorCode(err, serror.CodeDuplicateCard) {
				return errors.WithStack(err)
			}
		}

		chargeRequest := command.PaymentList{
			List: []*command.Payment{
				{
					ReturnGiftID:         returnGift.ID,
					ReturnGiftSnapshotID: int(returnGift.LatestSnapshotID.Int64),
					Amount:               1,
				},
			},
			Body: comment.Body,
		}

		if _, err := s.ChargeCommandService.Capture(user, &chargeRequest); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
