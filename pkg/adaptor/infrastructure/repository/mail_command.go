package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"go.uber.org/zap"
)

type (
	MailCommandRepositoryImpl struct {
		AWSSession *session.Session
		AWSConfig  config.AWS
	}

	MailCommandRepositoryForLocalImpl struct {
	}
)

func (r *MailCommandRepositoryImpl) SendTemplateMail(toEmail string, template entity.MailTemplate) error {
	data, err := template.ToJSON()
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	defaultData, err := template.DefaultData()
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	destinations := []*ses.BulkEmailDestination{
		{
			Destination: &ses.Destination{
				ToAddresses: aws.StringSlice([]string{toEmail}),
			},
			ReplacementTemplateData: aws.String(data),
		},
	}

	svc := ses.New(r.AWSSession)
	input := &ses.SendBulkTemplatedEmailInput{
		DefaultTemplateData: aws.String(defaultData),
		Destinations:        destinations,
		Source:              aws.String(r.AWSConfig.FromEmail),
		Template:            aws.String(template.TemplateName().String()),
	}

	_, err = svc.SendBulkTemplatedEmail(input)
	if err != nil {
		return errors.Wrap(err, "failed send template email by ses")
	}

	return nil
}

func (r *MailCommandRepositoryForLocalImpl) SendTemplateMail(toEmail string, template entity.MailTemplate) error {
	data, err := template.ToJSON()
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	logger.Info("Email", zap.String("toEmail", toEmail), zap.String("template", template.TemplateName().String()), zap.String("templateData", data))
	return nil
}
