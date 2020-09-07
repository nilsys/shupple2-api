package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Undefined, Male, Female)
*/
type Gender int

/*
ENUM(NEW = 1, RANKING, RECOMMEND)
*/
type MediaSortBy int

/*
ENUM(NEW = 1, RECOMMEND)
*/
type ReviewSortBy int

/*
ENUM(RANKING = 1, RECOMMEND)
*/
type UserSortBy int

/*
ENUM(Area = 1, SubArea, SubSubArea, TouristSpot, HashTag, User)
*/
type SuggestionType int

/*
ENUM(BUSINESS = 1, COUPLE, FAMILY, FRIEND, ONLY, WITHCHILD)
*/
type AccompanyingType int

/*
ENUM(Undefined, Japan, World)
*/
type AreaGroup int

/*
ENUM(Undefined, Area, SubArea, SubSubArea)
sub_sub_areaの子カテとその子孫がundefinedに分類される
*/
type AreaCategoryType int

/*
ENUM(Undefined, Theme = 1, SubTheme)
sub_themeの子カテとその子孫がundefinedに分類される
*/
type ThemeCategoryType int

/*
ENUM(Undefined, SpotCategory = 1, SubSpotCategory)
*/
type SpotCategoryType int

/*
ENUM(Undefined, Style, Scene, Gourmet, LifeStyle, Activity, Sport)
*/
type InterestGroup int

/*
ENUM(FAVORITE = 1, COMMENT, REPLY, FOLLOW, TAGGED)
*/
type NoticeActionType int

/*
ENUM(POST = 1, VLOG, REVIEW, COMMENT, REPLY, USER, COMIC)
*/
type NoticeActionTargetType int

/*
ENUM(Transaction = 1)
*/
type ContextKey int

/*
ENUM(REVIEW = 1, COMMENT, REPLY)
*/
type ReportTargetType int

/*
ENUM(UNKNOWN =1, SEXUAL, INAPPROPRIATE, COPYRIGHT, SELFHARM, LIE, UNRELATED, AD)
*/
type ReportReasonType int

/*
ENUM(New, LargeAmount, Push, Attention)
*/
type CfProjectSortBy int

/*
ENUM(ReservedTicket = 1, Other)
*/
type CfReturnGiftType int

/*
ENUM(ThanksPurchase = 1, ThanksPurchaseForNonLoginUser, ThanksPurchaseForOwner, DepositRequestForStayway, CfReturnGiftShippingNotification, ReserveRequestForOwner, ReserveRequestForUser, CfProjectAchievementNoticeForSupporter, CfProjectPostNewReportNoticeForSupporter)
*/
type MailTemplateName int

/*
ENUM(Undefined, OwnerUnconfirmed, OwnerConfirmed, Canceled)
宿泊券の場合はUndefined
*/
type PaymentCfReturnGiftOtherTypeStatus int

/*
ENUM(Undefined, Unreserved, Reserved, Canceled, Expired)
宿泊券以外の場合はUndefined
*/
type PaymentCfReturnGiftReservedTicketTypeStatus int

/*
ENUM(Common = 1, WP, CfProjectOwner, CfProjectAdmin)
Userの属性,Commonがデフォルト値で一般ユーザーを表す
CfProjectAdminは一般的にStaywayのユーザーが使用するAdmin権限
*/
type UserAttribute int

/*
ENUM(Available = 1, Unavailable, Done)
Available = 入金可能になった
Unavailable = 入金不可能になった
Done = 入金申請済になった
*/
type UserSalesReasonType int
