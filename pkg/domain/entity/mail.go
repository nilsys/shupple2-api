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

	ThanksPurchaseForNonLoginUserTemplate struct {
		ThanksPurchaseTemplate
		LoginURL string `json:"loginurl"`
	}

	ReserveRequestTemplateForOwnerTemplate struct {
		UserFullName            string `json:"userfullname"`
		UserFullNameKana        string `json:"userfullnamekana"`
		UserEmail               string `json:"useremail"`
		UserPhoneNumber         string `json:"userphonenumber"`
		ChargeID                string `json:"chargeid"`
		InquiryCode             string `json:"inquirycode"`
		CfReturnGiftTitle       string `json:"cfreturngifttitle"`
		CfReturnGiftDescription string `json:"cfreturngiftdesc"`
		Checkin                 string `json:"checkin"`
		Checkout                string `json:"checkout"`
		StayDays                string `json:"staydays"`
		AdultMemberCount        string `json:"adultmembercount"`
		ChildMemberCount        string `json:"childmembercount"`
		Remark                  string `json:"remark"`
	}

	ReserveRequestTemplateForUserTemplate struct {
		UserFullName            string `json:"userfullname"`
		UserFullNameKana        string `json:"userfullnamekana"`
		UserEmail               string `json:"useremail"`
		UserPhoneNumber         string `json:"userphonenumber"`
		ChargeID                string `json:"chargeid"`
		InquiryCode             string `json:"inquirycode"`
		CfReturnGiftTitle       string `json:"cfreturngifttitle"`
		CfReturnGiftDescription string `json:"cfreturngiftdesc"`
		Checkin                 string `json:"checkin"`
		Checkout                string `json:"checkout"`
		StayDays                string `json:"staydays"`
		AdultMemberCount        string `json:"adultmembercount"`
		ChildMemberCount        string `json:"childmembercount"`
		Remark                  string `json:"remark"`
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

func NewThanksPurchaseForNonLoginUserTemplate(ownerName, returnGiftDesc, chargeID, systemFee, price, userEmail, userShippingAddress, userTel, userName, loginURL string) *ThanksPurchaseForNonLoginUserTemplate {
	return &ThanksPurchaseForNonLoginUserTemplate{
		ThanksPurchaseTemplate: *NewThanksPurchaseTemplate(ownerName, returnGiftDesc, chargeID, systemFee, price, userEmail, userShippingAddress, userTel, userName),
		LoginURL:               loginURL,
	}
}

func NewReserveRequestForOwnerTemplate(fullName, fullNameKana, email, phonenum, chargeID, inquiryCode, giftTitle, giftDesc, checkin, checkout, staydays, adultcount, childcount, remark string) *ReserveRequestTemplateForOwnerTemplate {
	return &ReserveRequestTemplateForOwnerTemplate{
		UserFullName:            fullName,
		UserFullNameKana:        fullNameKana,
		UserEmail:               email,
		UserPhoneNumber:         phonenum,
		ChargeID:                chargeID,
		InquiryCode:             inquiryCode,
		CfReturnGiftTitle:       giftTitle,
		CfReturnGiftDescription: giftDesc,
		Checkin:                 checkin,
		Checkout:                checkout,
		StayDays:                staydays,
		AdultMemberCount:        adultcount,
		ChildMemberCount:        childcount,
		Remark:                  remark,
	}
}

func NewReserveRequestForUserTemplate(fullName, fullNameKana, email, phonenum, chargeID, inquiryCode, giftTitle, giftDesc, checkin, checkout, staydays, adultcount, childcount, remark string) *ReserveRequestTemplateForUserTemplate {
	return &ReserveRequestTemplateForUserTemplate{
		UserFullName:            fullName,
		UserFullNameKana:        fullNameKana,
		UserEmail:               email,
		UserPhoneNumber:         phonenum,
		ChargeID:                chargeID,
		InquiryCode:             inquiryCode,
		CfReturnGiftTitle:       giftTitle,
		CfReturnGiftDescription: giftDesc,
		Checkin:                 checkin,
		Checkout:                checkout,
		StayDays:                staydays,
		AdultMemberCount:        adultcount,
		ChildMemberCount:        childcount,
		Remark:                  remark,
	}
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

func (t *ThanksPurchaseForNonLoginUserTemplate) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameThanksPurchaseForNonLoginUser
}

func (t *ThanksPurchaseForNonLoginUserTemplate) DefaultData() (string, error) {
	s := ThanksPurchaseForNonLoginUserTemplate{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *ThanksPurchaseForNonLoginUserTemplate) ToJSON() (string, error) {
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

func (t *ReserveRequestTemplateForUserTemplate) TemplateName() model.MailTemplateName {
	return model.MailTemplateNameReserveRequestForUser
}

func (t *ReserveRequestTemplateForUserTemplate) DefaultData() (string, error) {
	s := ReserveRequestTemplateForUserTemplate{}
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}

func (t *ReserveRequestTemplateForUserTemplate) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed marshal")
	}
	return string(bytes), nil
}
