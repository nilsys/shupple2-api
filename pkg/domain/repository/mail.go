package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	MailCommandRepository interface {
		SendTemplateMail(toEmail string, template entity.MailTemplate) error
	}
)
