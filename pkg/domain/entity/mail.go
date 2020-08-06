package entity

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	MailTemplate interface {
		TemplateName() model.MailTemplateName
		DefaultData() (string, error)
		ToJSON() (string, error)
	}

	ThanksPurchaseTemplate struct {
		OwnerName             string `json:"ownername"`
		ReturnGiftDescription string `json:"returngiftdescription"`
		ChargeID              string `json:"chargeid"`
		SystemFee             string `json:"systemfee"`
		Price                 string `json:"price"`
		UserEmail             string `json:"useremail"`
		UserShippingAddress   string `json:"usershippingaddress"`
		UserTel               string `json:"usertel"`
		UserName              string `json:"username"`
	}

	ReserveRequestTemplateForOwnerTemplate struct {
		UserFullName          string `json:"userfullname"`
		UserFullNameKana      string `json:"userfullnamekana"`
		UserEmail             string `json:"useremail"`
		UserPhoneNumber       string `json:"userphonenumber"`
		ChargeID              string `json:"chargeid"`
		ReturnGiftDescription string `json:"returngiftdescription"`
		Checkin               string `json:"checkin"`
		Checkout              string `json:"checkout"`
		StayDays              string `json:"staydays"`
		AdultMemberCount      string `json:"adultmembercount"`
		ChildMemberCount      string `json:"childmembercount"`
	}

	CfProjectAchievementNoticeForSupporter struct {
		ProjectID         string `json:"projectid"`
		ProjectTitle      string `json:"projecttitle"`
		ProjectOwnerEmail string `json:"projectowneremail"`
	}

	CfProjectPostNewReportNoticeForSupporter struct {
		ProjectID    string `json:"projectid"`
		ProjectTitle string `json:"projecttitle"`
		PostTitle    string `json:"posttitle"`
		PostSlug     string `json:"postslug"`
		PostBody     string `json:"postbody"`
	}
)

func NewThanksPurchaseTemplate(ownerName, returnGiftDesc, chargeID, systemFee, price, userEmail, userShippingAddress, userTel, userName string) *ThanksPurchaseTemplate {
	return &ThanksPurchaseTemplate{
		OwnerName:             ownerName,
		ReturnGiftDescription: returnGiftDesc,
		ChargeID:              chargeID,
		SystemFee:             systemFee,
		Price:                 price,
		UserEmail:             userEmail,
		UserShippingAddress:   userShippingAddress,
		UserTel:               userTel,
		UserName:              userName,
	}
}

func NewReserveRequestTemplate(fullName, fullNameKana, email, phonenum, chargeID, giftDesc, checkin, checkout, staydays, adultcount, childcount string) *ReserveRequestTemplateForOwnerTemplate {
	return &ReserveRequestTemplateForOwnerTemplate{
		UserFullName:          fullName,
		UserFullNameKana:      fullNameKana,
		UserEmail:             email,
		UserPhoneNumber:       phonenum,
		ChargeID:              chargeID,
		ReturnGiftDescription: giftDesc,
		Checkin:               checkin,
		Checkout:              checkout,
		StayDays:              staydays,
		AdultMemberCount:      adultcount,
		ChildMemberCount:      childcount,
	}
}

func NewReserveRequestTemplateFromCfReserveRequest(req *CfReserveRequest, chargeID, giftDesc string) *ReserveRequestTemplateForOwnerTemplate {
	return NewReserveRequestTemplate(req.FullNameMailFmt(), req.FullNameKanaMailFmt(), req.Email, req.PhoneNumber, chargeID, giftDesc, req.Checkin, req.Checkout, strconv.Itoa(req.StayDays), strconv.Itoa(req.AdultMemberCount), strconv.Itoa(req.ChildMemberCount))
}

func NewCfProjectAchievementNoticeForSupporter(cfProjectID int, cfProjectTitle, cfProjectOwnerEmail string) *CfProjectAchievementNoticeForSupporter {
	return &CfProjectAchievementNoticeForSupporter{
		ProjectID:         strconv.Itoa(cfProjectID),
		ProjectTitle:      cfProjectTitle,
		ProjectOwnerEmail: cfProjectOwnerEmail,
	}
}

func NewCfProjectPostNewReportNoticeForSupporter(projectID int, projectTitle, postTitle, postSlug, postBody string) *CfProjectPostNewReportNoticeForSupporter {
	return &CfProjectPostNewReportNoticeForSupporter{
		ProjectID:    strconv.Itoa(projectID),
		ProjectTitle: projectTitle,
		PostTitle:    postTitle,
		PostSlug:     postSlug,
		PostBody:     postBody,
	}
}

func (t *ThanksPurchaseTemplate) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameThanksPurchase
}

func (t *ThanksPurchaseTemplate) DefaultData() (string, error) {
	s := ThanksPurchaseTemplate{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *ThanksPurchaseTemplate) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *ReserveRequestTemplateForOwnerTemplate) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameReserveRequestForOwner
}

func (t *ReserveRequestTemplateForOwnerTemplate) DefaultData() (string, error) {
	s := ReserveRequestTemplateForOwnerTemplate{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *ReserveRequestTemplateForOwnerTemplate) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *CfProjectAchievementNoticeForSupporter) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameCfProjectAchievementNoticeForSupporter
}

func (t *CfProjectAchievementNoticeForSupporter) DefaultData() (string, error) {
	s := CfProjectAchievementNoticeForSupporter{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *CfProjectAchievementNoticeForSupporter) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *CfProjectPostNewReportNoticeForSupporter) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameCfProjectPostNewReportNoticeForSupporter
}

func (t *CfProjectPostNewReportNoticeForSupporter) DefaultData() (string, error) {
	s := CfProjectPostNewReportNoticeForSupporter{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *CfProjectPostNewReportNoticeForSupporter) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}
